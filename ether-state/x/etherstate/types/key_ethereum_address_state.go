package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EthereumAddressStateKeyPrefix is the prefix to retrieve all EthereumAddressState
	EthereumAddressStateKeyPrefix = "EthereumAddressState/value/"
)

// EthereumAddressStateKey returns the store key to retrieve a EthereumAddressState from the index fields
func EthereumAddressStateKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
