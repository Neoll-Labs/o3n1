package keeper

import (
	"statether/x/statether/types"
)

var _ types.QueryServer = Keeper{}
