package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

// SetEthereumAddressState set a specific ethereumAddressState in the store from its index
func (k Keeper) SetEthereumAddressState(ctx sdk.Context, ethereumAddressState types.EthereumAddressState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressStateKeyPrefix))
	b := k.cdc.MustMarshal(&ethereumAddressState)
	store.Set(types.EthereumAddressStateKey(
		ethereumAddressState.Index,
	), b)
}

// GetEthereumAddressState returns a ethereumAddressState from its index
func (k Keeper) GetEthereumAddressState(
	ctx sdk.Context,
	index string,

) (val types.EthereumAddressState, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressStateKeyPrefix))

	b := store.Get(types.EthereumAddressStateKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEthereumAddressState removes a ethereumAddressState from the store
func (k Keeper) RemoveEthereumAddressState(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressStateKeyPrefix))
	store.Delete(types.EthereumAddressStateKey(
		index,
	))
}

// GetAllEthereumAddressState returns all ethereumAddressState
func (k Keeper) GetAllEthereumAddressState(ctx sdk.Context) (list []types.EthereumAddressState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressStateKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EthereumAddressState
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
