package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EthAddressKeyPrefix is the prefix to retrieve all EthAddress
	EthAddressKeyPrefix = "EthAddress/value/"
)

// EthAddressKey returns the store key to retrieve a EthAddress from the index fields
func EthAddressKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
