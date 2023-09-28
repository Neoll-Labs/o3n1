package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		EthereumAddressList:      []EthereumAddress{},
		EthereumAddressStateList: []EthereumAddressState{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in ethereumAddress
	ethereumAddressIndexMap := make(map[string]struct{})

	for _, elem := range gs.EthereumAddressList {
		index := string(EthereumAddressKey(elem.Index))
		if _, ok := ethereumAddressIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ethereumAddress")
		}
		ethereumAddressIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in ethereumAddressState
	ethereumAddressStateIndexMap := make(map[string]struct{})

	for _, elem := range gs.EthereumAddressStateList {
		index := string(EthereumAddressStateKey(elem.Index))
		if _, ok := ethereumAddressStateIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for ethereumAddressState")
		}
		ethereumAddressStateIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
