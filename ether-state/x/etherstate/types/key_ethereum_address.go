package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// EthereumAddressKeyPrefix is the prefix to retrieve all EthereumAddress
	EthereumAddressKeyPrefix = "EthereumAddress/value/"
)

// EthereumAddressKey returns the store key to retrieve a EthereumAddress from the index fields
func EthereumAddressKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
