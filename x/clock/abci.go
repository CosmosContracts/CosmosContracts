package clock

import (
	"log"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmosContracts/juno/v18/x/clock/keeper"
	"github.com/CosmosContracts/juno/v18/x/clock/types"
)

// EndBlocker executes on contracts at the end of the block.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	message := []byte(types.EndBlockSudoMessage)

	p := k.GetParams(ctx)

	contracts := k.GetAllContracts(ctx, false)
	errorExecs := make([]string, len(contracts))
	errorExists := false

	for idx, addr := range contracts {
		contract, err := sdk.AccAddressFromBech32(addr)
		if err != nil {
			errorExecs[idx] = addr
			continue
		}

		childCtx := ctx.WithGasMeter(sdk.NewGasMeter(p.ContractGasLimit))
		_, err = k.GetContractKeeper().Sudo(childCtx, contract, message)
		if err != nil {
			errorExists = true
			errorExecs[idx] = addr

			// Jail contract on error
			err = k.JailContract(ctx, addr)
			if err != nil {
				log.Printf("[x/clock] Failed to Error Contract %s: %v", addr, err)
			}
		}
	}

	// Log errors if present
	if errorExists {
		log.Printf("[x/clock] Execute Errors: %v", errorExecs)
	}
}
