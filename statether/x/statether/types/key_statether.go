package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// StatetherKeyPrefix is the prefix to retrieve all Statether
	StatetherKeyPrefix = "Statether/value/"
)

// StatetherKey returns the store key to retrieve a Statether from the index fields
func StatetherKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
