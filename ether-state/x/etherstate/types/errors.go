package types

// DONTCOVER

import (
	"cosmossdk.io/errors"
)

// x/statether module sentinel errors
var (
	ErrInvalidEthereumAddress                = errors.Register(ModuleName, 1101, "invalid ethereum address")
	ErrInvalidNonce                          = errors.Register(ModuleName, 1102, "invalid nonce")
	ErrInvalidBlockNumber                    = errors.Register(ModuleName, 1103, "invalid block number")
	ErrInvalidStoragePosition                = errors.Register(ModuleName, 1104, "invalid storage position")
	ErrEthereumAddressAlreadyOnRequiredState = errors.Register(ModuleName, 1105, "ethereum address already in required state")
)
