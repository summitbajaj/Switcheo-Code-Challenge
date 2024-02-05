package keeper

import (
	"context"
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
	"match/x/match/types"
)

// GetMatchInfoCount get the total number of matchInfo
func (k Keeper) GetMatchInfoCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MatchInfoCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetMatchInfoCount set the total number of matchInfo
func (k Keeper) SetMatchInfoCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MatchInfoCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendMatchInfo appends a matchInfo in the store with a new id and update the count
func (k Keeper) AppendMatchInfo(
	ctx context.Context,
	matchInfo types.MatchInfo,
) uint64 {
	// Create the matchInfo
	count := k.GetMatchInfoCount(ctx)

	// Set the ID of the appended value
	matchInfo.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MatchInfoKey))
	appendedValue := k.cdc.MustMarshal(&matchInfo)
	store.Set(GetMatchInfoIDBytes(matchInfo.Id), appendedValue)

	// Update matchInfo count
	k.SetMatchInfoCount(ctx, count+1)

	return count
}

// SetMatchInfo set a specific matchInfo in the store
func (k Keeper) SetMatchInfo(ctx context.Context, matchInfo types.MatchInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MatchInfoKey))
	b := k.cdc.MustMarshal(&matchInfo)
	store.Set(GetMatchInfoIDBytes(matchInfo.Id), b)
}

// GetMatchInfo returns a matchInfo from its id
func (k Keeper) GetMatchInfo(ctx context.Context, id uint64) (val types.MatchInfo, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MatchInfoKey))
	b := store.Get(GetMatchInfoIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveMatchInfo removes a matchInfo from the store
func (k Keeper) RemoveMatchInfo(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MatchInfoKey))
	store.Delete(GetMatchInfoIDBytes(id))
}

// GetAllMatchInfo returns all matchInfo
func (k Keeper) GetAllMatchInfo(ctx context.Context) (list []types.MatchInfo) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MatchInfoKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.MatchInfo
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMatchInfoIDBytes returns the byte representation of the ID
func GetMatchInfoIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.MatchInfoKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
