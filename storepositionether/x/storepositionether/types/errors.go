package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/storepositionether module sentinel errors
var (
	ErrInvalidEthAddress = errors.Register(ModuleName, 1101, "invalid eth address")
)
