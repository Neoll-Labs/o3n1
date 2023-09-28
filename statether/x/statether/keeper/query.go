package keeper

import (
	"github.com/nelsonstr/o3n1/statether/x/statether/types"
)

var _ types.QueryServer = Keeper{}
