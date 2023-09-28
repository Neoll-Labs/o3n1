package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
)

// SetStatether set a specific statether in the store from its index
func (k Keeper) SetStatether(ctx sdk.Context, statether types.Statether) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatetherKeyPrefix))
	b := k.cdc.MustMarshal(&statether)
	store.Set(types.StatetherKey(
		statether.Index,
	), b)
}

// GetStatether returns a statether from its index
func (k Keeper) GetStatether(
	ctx sdk.Context,
	index string,

) (val types.Statether, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatetherKeyPrefix))

	b := store.Get(types.StatetherKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveStatether removes a statether from the store
func (k Keeper) RemoveStatether(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatetherKeyPrefix))
	store.Delete(types.StatetherKey(
		index,
	))
}

// GetAllStatether returns all statether
func (k Keeper) GetAllStatether(ctx sdk.Context) (list []types.Statether) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.StatetherKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Statether
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
