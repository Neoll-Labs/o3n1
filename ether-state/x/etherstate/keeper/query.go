package keeper

import (
	"ether-state/x/etherstate/types"
)

var _ types.QueryServer = Keeper{}
