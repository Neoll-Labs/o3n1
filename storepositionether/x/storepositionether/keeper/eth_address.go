package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"storepositionether/x/storepositionether/types"
)

// SetEthAddress set a specific ethAddress in the store from its index
func (k Keeper) SetEthAddress(ctx sdk.Context, ethAddress types.EthAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthAddressKeyPrefix))
	b := k.cdc.MustMarshal(&ethAddress)
	store.Set(types.EthAddressKey(
		ethAddress.Index,
	), b)
}

// GetEthAddress returns a ethAddress from its index
func (k Keeper) GetEthAddress(
	ctx sdk.Context,
	index string,

) (val types.EthAddress, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthAddressKeyPrefix))

	b := store.Get(types.EthAddressKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEthAddress removes a ethAddress from the store
func (k Keeper) RemoveEthAddress(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthAddressKeyPrefix))
	store.Delete(types.EthAddressKey(
		index,
	))
}

// GetAllEthAddress returns all ethAddress
func (k Keeper) GetAllEthAddress(ctx sdk.Context) (list []types.EthAddress) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EthAddressKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.EthAddress
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
