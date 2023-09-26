package keeper

import (
	"demo/x/store/types"
)

var _ types.QueryServer = Keeper{}
