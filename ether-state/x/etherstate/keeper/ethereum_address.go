package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

// SetEthereumAddress set a specific ethereumAddress in the store from its index
func (k Keeper) SetEthereumAddress(ctx sdk.Context, ethereumAddress types.EthereumAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressKeyPrefix))
	b := k.cdc.MustMarshal(&ethereumAddress)
	store.Set(types.EthereumAddressKey(
		ethereumAddress.Index,
	), b)
}

// GetEthereumAddress returns a ethereumAddress from its index
func (k Keeper) GetEthereumAddress(
	ctx sdk.Context,
	index string,

) (val types.EthereumAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressKeyPrefix))

	b := store.Get(types.EthereumAddressKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEthereumAddress removes a ethereumAddress from the store
func (k Keeper) RemoveEthereumAddress(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressKeyPrefix))
	store.Delete(types.EthereumAddressKey(
		index,
	))
}

// GetAllEthereumAddress returns all ethereumAddress
func (k Keeper) GetAllEthereumAddress(ctx sdk.Context) (list []types.EthereumAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthereumAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EthereumAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
