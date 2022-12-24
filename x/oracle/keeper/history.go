package keeper

import (
	"errors"
	"time"

	"github.com/CosmosContracts/juno/v12/x/oracle/types"
	"github.com/CosmosContracts/juno/v12/x/oracle/util"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// storeHistorical data writes to the store, in all needed indexing.
// TODO : testing
func (k Keeper) storeHistoricalData(ctx sdk.Context, denom string, entry types.PriceHistoryEntry) {
	store := ctx.KVStore(k.storeKey)
	key := types.FormatHistoricalDenomIndexKey(entry.PriceUpdateTime, denom)

	bz, err := k.cdc.Marshal(&entry)
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

// getHistoryEntryAtOrBeforeTime on a given input (denom, t)
// returns the PriceHistoryEntry from state for (denom, t'),
// where t' is such that:
// * t' <= t
// * there exists no `t” <= t` in state, where `t' < t”`
// TODO : testing
func (k Keeper) getHistoryEntryAtOrBeforeTime(ctx sdk.Context, denom string, t time.Time) (types.PriceHistoryEntry, error) {
	store := ctx.KVStore(k.storeKey)
	// reverseIterator not catch end key => Need this scope to catch if the value is in end key
	key := types.FormatHistoricalDenomIndexKey(t, denom)
	bz := store.Get(key)
	if bz != nil {
		var entry types.PriceHistoryEntry
		err := k.cdc.Unmarshal(bz, &entry)
		if err != nil {
			return types.PriceHistoryEntry{}, err
		}
		return entry, nil
	}

	startKey := types.FormatHistoricalDenomIndexPrefix(denom)
	endKey := types.FormatHistoricalDenomIndexKey(t, denom)
	reverseIterate := true

	entry, err := util.GetFirstValueInRange(store, startKey, endKey, reverseIterate, k.ParseTwapFromBz)
	if err != nil {
		return types.PriceHistoryEntry{}, err
	}

	return entry, nil
}

// getHistoryEntryAtOrAfterTime on a given input (denom, t)
// returns the PriceHistoryEntry from state for (denom, t'),
// where t' is such that:
// * t' => t
// * there exists no `t” => t` in state, where `t' > t”`
// TODO : testing
func (k Keeper) getHistoryEntryAtOrAfterTime(ctx sdk.Context, denom string, t time.Time) (types.PriceHistoryEntry, error) {
	store := ctx.KVStore(k.storeKey)

	startKey := types.FormatHistoricalDenomIndexKey(t, denom)
	reverseIterate := false

	entry, err := util.GetFirstValueInRange(store, startKey, nil, reverseIterate, k.ParseTwapFromBz)
	if err != nil {
		return types.PriceHistoryEntry{}, err
	}

	return entry, nil
}

// getHistoryEntryBetweenTime on a given input (denom, t)
// returns the PriceHistoryEntry from state for (denom, t'),
// TODO : testing
func (k Keeper) getHistoryEntryBetweenTime(ctx sdk.Context, denom string, start time.Time, end time.Time) (types.PriceHistoryEntry, error) { //nolint:unused
	store := ctx.KVStore(k.storeKey)

	startKey := types.FormatHistoricalDenomIndexKey(start, denom)
	endKey := types.FormatHistoricalDenomIndexKey(end, denom)

	reverseIterate := true

	entry, err := util.GetFirstValueInRange(store, startKey, endKey, reverseIterate, k.ParseTwapFromBz)
	if err != nil {
		return types.PriceHistoryEntry{}, err
	}

	return entry, nil
}

func (k Keeper) ParseTwapFromBz(bz []byte) (entry types.PriceHistoryEntry, err error) {
	if len(bz) == 0 {
		return types.PriceHistoryEntry{}, errors.New("history entry not found")
	}
	err = k.cdc.Unmarshal(bz, &entry)
	return entry, err
}

// TODO : testing
func (k Keeper) RemoveHistoryEntryAtOrBeforeTime(ctx sdk.Context, denom string, t time.Time) {
	store := ctx.KVStore(k.storeKey)

	startKey := types.FormatHistoricalDenomIndexPrefix(denom)
	endKey := types.FormatHistoricalDenomIndexKey(t, denom)
	reverseIterate := true

	_ = util.RemoveValueInRange(store, startKey, endKey, reverseIterate)
}

// TODO : testing
func (k Keeper) SetPriceHistoryEntry(ctx sdk.Context, denom string, t time.Time, exchangeRate sdk.Dec, votingPeriodCount uint64) {
	entry := types.PriceHistoryEntry{
		Price:           exchangeRate,
		VotePeriodCount: votingPeriodCount,
		PriceUpdateTime: t,
	}

	k.storeHistoricalData(ctx, denom, entry)
}
