package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		StatetherList: []Statether{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in statether
	statetherIndexMap := make(map[string]struct{})

	for _, elem := range gs.StatetherList {
		index := string(StatetherKey(elem.Index))
		if _, ok := statetherIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for statether")
		}
		statetherIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
