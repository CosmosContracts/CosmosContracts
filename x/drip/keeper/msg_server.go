package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/CosmosContracts/juno/v15/x/drip/types"
)

var _ types.MsgServer = &Keeper{}

// DistributeTokens distribute tokens to all stakers at the next block
// TODO: Impl
func (k Keeper) DistributeTokens(
	goCtx context.Context,
	msg *types.MsgDistributeTokens,
) (*types.MsgDistributeTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	params := k.GetParams(ctx)
	if !params.EnableDrip {
		return nil, types.ErrDripDisabled
	}

	// Check if sender is allowed
	authorized := false
	for _, addr := range params.AllowedAddresses {
		if msg.SenderAddress == addr {
			authorized = true
			break
		}
	}

	if !authorized {
		return nil, types.ErrDripNotAllowed
	}

	// Get sender
	sender, err := sdk.AccAddressFromBech32(msg.SenderAddress)
	if err != nil {
		return nil, err
	}

	if err := k.SendCoinsFromAccountToFeeCollector(ctx, sender, msg.Amount); err != nil {
		return nil, err
	}

	return &types.MsgDistributeTokensResponse{}, nil
}
