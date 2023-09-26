package keeper

import (
	"demo/x/demo/types"
)

var _ types.QueryServer = Keeper{}
