package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/storepositionether module sentinel errors
var (
	ErrInvalidEthAddress      = errors.Register(ModuleName, 1101, "invalid eth address")
	ErrInvalidNonce           = errors.Register(ModuleName, 1102, "invalid nonce")
	ErrInvalidBlock           = errors.Register(ModuleName, 1103, "invalid block number")
	ErrInvalidStoragePosition = errors.Register(ModuleName, 1104, "invalid storage position")
)
