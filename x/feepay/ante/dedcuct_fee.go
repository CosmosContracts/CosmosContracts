package ante

import (
	"fmt"
	"math"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	feepaykeeper "github.com/CosmosContracts/juno/v17/x/feepay/keeper"
	feepaytypes "github.com/CosmosContracts/juno/v17/x/feepay/types"
	globalfeekeeper "github.com/CosmosContracts/juno/v17/x/globalfee/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
)

// DeductFeeDecorator deducts fees from the first signer of the tx
// If the first signer does not have the funds to pay for the fees, return with InsufficientFunds error
// Call next AnteHandler if fees successfully deducted
// CONTRACT: Tx must implement FeeTx interface to use DeductFeeDecorator
//
// Additionally, the Deduct Fee ante is a fork of the SDK's DeductFeeDecorator. This decorator looks for single
// message transactions with no provided fee. If they correspond to a registered FeePay Contract, the FeePay
// module will cover the cost of the fee (if the balance permits).
type DeductFeeDecorator struct {
	feepayKeeper    feepaykeeper.Keeper
	globalfeeKeeper globalfeekeeper.Keeper
	accountKeeper   ante.AccountKeeper
	bankKeeper      bankkeeper.Keeper
	feegrantKeeper  ante.FeegrantKeeper
	// TxFeeChecker check if the provided fee is enough and returns the effective fee and tx priority,
	// the effective fee should be deducted later, and the priority should be returned in abci response.
	// type TxFeeChecker func(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error)
	txFeeChecker ante.TxFeeChecker

	bondDenom string
}

func NewDeductFeeDecorator(fpk feepaykeeper.Keeper, gfk globalfeekeeper.Keeper, ak ante.AccountKeeper, bk bankkeeper.Keeper, fgk ante.FeegrantKeeper, tfc ante.TxFeeChecker, bondDenom string) DeductFeeDecorator {
	if tfc == nil {
		tfc = checkTxFeeWithValidatorMinGasPrices
	}

	return DeductFeeDecorator{
		feepayKeeper:    fpk,
		globalfeeKeeper: gfk,
		accountKeeper:   ak,
		bankKeeper:      bk,
		feegrantKeeper:  fgk,
		txFeeChecker:    tfc,
		bondDenom:       bondDenom,
	}
}

func (dfd DeductFeeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		// return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
		return ctx, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if !simulate && ctx.BlockHeight() > 0 && feeTx.GetGas() == 0 {
		return ctx, errorsmod.Wrap(sdkerrors.ErrInvalidGasLimit, "must provide positive gas")
	}

	var (
		priority int64
		err      error
	)

	fee := feeTx.GetFee()
	if !simulate {
		fee, priority, err = dfd.txFeeChecker(ctx, tx)
		if err != nil {
			return ctx, err
		}
	}
	if err := dfd.checkDeductFee(ctx, tx, fee); err != nil {
		return ctx, err
	}

	newCtx := ctx.WithPriority(priority)

	return next(newCtx, tx, simulate)
}

func (dfd DeductFeeDecorator) checkDeductFee(ctx sdk.Context, sdkTx sdk.Tx, fee sdk.Coins) error {
	feeTx, ok := sdkTx.(sdk.FeeTx)
	if !ok {
		return errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	if addr := dfd.accountKeeper.GetModuleAddress(types.FeeCollectorName); addr == nil {
		return fmt.Errorf("fee collector module account (%s) has not been set", types.FeeCollectorName)
	}

	feePayer := feeTx.FeePayer()
	feeGranter := feeTx.FeeGranter()
	deductFeesFrom := feePayer

	// if feegranter set deduct fee from feegranter account.
	// this works with only when feegrant enabled.
	if feeGranter != nil {
		if &dfd.feegrantKeeper == nil {
			return sdkerrors.ErrInvalidRequest.Wrap("fee grants are not enabled")
		} else if !feeGranter.Equals(feePayer) {
			err := dfd.feegrantKeeper.UseGrantedFees(ctx, feeGranter, feePayer, fee, sdkTx.GetMsgs())
			if err != nil {
				return errorsmod.Wrapf(err, "%s does not allow to pay fees for %s", feeGranter, feePayer)
			}
		}

		deductFeesFrom = feeGranter
	}

	deductFeesFromAcc := dfd.accountKeeper.GetAccount(ctx, deductFeesFrom)
	if deductFeesFromAcc == nil {
		return sdkerrors.ErrUnknownAddress.Wrapf("fee payer address: %s does not exist", deductFeesFrom)
	}

	if isValidFeePayTransaction(ctx, sdkTx, fee) {
		err := dfd.handleZeroFees(ctx, deductFeesFromAcc, sdkTx, fee)
		if err != nil {
			return err
		}
	} else {
		// Std sdk route
		err := DeductFees(dfd.bankKeeper, ctx, deductFeesFromAcc, fee)
		if err != nil {
			return err
		}
	}

	events := sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeTx,
			sdk.NewAttribute(sdk.AttributeKeyFee, fee.String()),
			sdk.NewAttribute(sdk.AttributeKeyFeePayer, deductFeesFrom.String()),
		),
	}
	ctx.EventManager().EmitEvents(events)

	return nil
}

// Handle zero fee transactions for fee prepay module
func (dfd DeductFeeDecorator) handleZeroFees(ctx sdk.Context, deductFeesFromAcc types.AccountI, tx sdk.Tx, fee sdk.Coins) error {
	ctx.Logger().Error("HandleZeroFees", "Starting", true)

	msg := tx.GetMsgs()[0]
	cw := msg.(*wasmtypes.MsgExecuteContract)

	// We need to check if it is a valid contract. Utilize the FeePay Keeper for validation
	if !dfd.feepayKeeper.IsValidContract(ctx, cw.GetContract()) {
		return sdkerrors.ErrInvalidRequest.Wrapf("contract %s is not registered for fee pay", cw.GetContract())
	}

	// Get the fee price in the chain denom
	var feePrice sdk.DecCoin
	for _, c := range dfd.globalfeeKeeper.GetParams(ctx).MinimumGasPrices {
		if c.Denom == dfd.bondDenom {
			feePrice = c
		}
	}

	ctx.Logger().Error("HandleZeroFees", "FeePrice", feePrice)

	// Get the tx gas
	feeTx := tx.(sdk.FeeTx)
	gas := sdkmath.LegacyNewDec(int64(feeTx.GetGas()))

	ctx.Logger().Error("HandleZeroFees", "Gas", gas)

	requiredFee := feePrice.Amount.Mul(gas).Ceil().RoundInt()

	ctx.Logger().Error("HandleZeroFees", "RequiredFee", requiredFee)

	// Create an array of coins, storing the required fee
	payment := sdk.NewCoins(sdk.NewCoin(feePrice.Denom, requiredFee))

	ctx.Logger().Error("HandleZeroFees", "Payment", payment)

	// Cover the fees of the transaction, send from FeePay Module to FeeCollector Module
	err := dfd.bankKeeper.SendCoinsFromModuleToModule(ctx, feepaytypes.ModuleName, types.FeeCollectorName, payment)

	// Throw transfer errors
	if err != nil {
		ctx.Logger().Error("HandleZeroFees", "Error transfering funds from module to module", err)
		return sdkerrors.ErrInsufficientFunds.Wrapf("error transfering funds from module to module: %s", err)
	}

	ctx.Logger().Error("HandleZeroFees", "Ending", true)
	return nil
}

// DeductFees deducts fees from the given account.
func DeductFees(bankKeeper types.BankKeeper, ctx sdk.Context, acc types.AccountI, fees sdk.Coins) error {
	if !fees.IsValid() {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "invalid fee amount: %s", fees)
	}

	// TODO: if 0 fees are sent, then the module account needs to pay it. (prepay module) ELSE have the standard user
	err := bankKeeper.SendCoinsFromAccountToModule(ctx, acc.GetAddress(), types.FeeCollectorName, fees)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInsufficientFunds, err.Error())
	}

	return nil
}

// from the SDK pulled out
func checkTxFeeWithValidatorMinGasPrices(ctx sdk.Context, tx sdk.Tx) (sdk.Coins, int64, error) {
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return nil, 0, errorsmod.Wrap(sdkerrors.ErrTxDecode, "Tx must be a FeeTx")
	}

	feeCoins := feeTx.GetFee()
	gas := feeTx.GetGas()

	// Ensure that the provided fees meet a minimum threshold for the validator,
	// if this is a CheckTx. This is only for local mempool purposes, and thus
	// is only ran on check tx.
	if ctx.IsCheckTx() && !isValidFeePayTransaction(ctx, tx, feeTx.GetFee()) {
		minGasPrices := ctx.MinGasPrices()
		if !minGasPrices.IsZero() {
			requiredFees := make(sdk.Coins, len(minGasPrices))

			// Determine the required fees by multiplying each required minimum gas
			// price by the gas limit, where fee = ceil(minGasPrice * gasLimit).
			glDec := sdkmath.LegacyNewDec(int64(gas))
			for i, gp := range minGasPrices {
				fee := gp.Amount.Mul(glDec)
				requiredFees[i] = sdk.NewCoin(gp.Denom, fee.Ceil().RoundInt())
			}

			if !feeCoins.IsAnyGTE(requiredFees) {
				return nil, 0, errorsmod.Wrapf(sdkerrors.ErrInsufficientFee, "insufficient fees; got: %s required: %s", feeCoins, requiredFees)
			}
		}
	}

	priority := getTxPriority(feeCoins, int64(gas))
	return feeCoins, priority, nil
}

// getTxPriority returns a naive tx priority based on the amount of the smallest denomination of the gas price
// provided in a transaction.
// NOTE: This implementation should be used with a great consideration as it opens potential attack vectors
// where txs with multiple coins could not be prioritize as expected.
func getTxPriority(fee sdk.Coins, gas int64) int64 {
	var priority int64
	for _, c := range fee {
		p := int64(math.MaxInt64)
		gasPrice := c.Amount.QuoRaw(gas)
		if gasPrice.IsInt64() {
			p = gasPrice.Int64()
		}
		if priority == 0 || p < priority {
			priority = p
		}
	}

	return priority
}

// Check if a transaction should be processed as a FeePay transaction.
// A valid FeePay transaction has no fee, and only 1 message for executing a contract.
func isValidFeePayTransaction(ctx sdk.Context, tx sdk.Tx, fee sdk.Coins) bool {

	ctx.Logger().Error("FeePayAnte", "IsZero", fee.IsZero(), "Msgs", len(tx.GetMsgs()))

	// Check if fee is zero, and tx has only 1 message for executing a contract
	if fee.IsZero() && len(tx.GetMsgs()) == 1 {
		_, ok := (tx.GetMsgs()[0]).(*wasmtypes.MsgExecuteContract)

		ctx.Logger().Error("FeePayAnte", "IsCWExecuteContract", ok)

		return ok
	}

	// The transaction includes a fee, has more than 1 message, or
	// has a single message that is not for executing a contract
	return false
}
