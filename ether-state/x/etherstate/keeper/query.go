package keeper

import (
	"github.com/nelsonstr/o3n1/ether-state/x/etherstate/types"
)

var _ types.QueryServer = Keeper{}
