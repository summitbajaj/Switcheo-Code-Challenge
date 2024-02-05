package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		MatchInfoList: []MatchInfo{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in matchInfo
	matchInfoIdMap := make(map[uint64]bool)
	matchInfoCount := gs.GetMatchInfoCount()
	for _, elem := range gs.MatchInfoList {
		if _, ok := matchInfoIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for matchInfo")
		}
		if elem.Id >= matchInfoCount {
			return fmt.Errorf("matchInfo id should be lower or equal than the last id")
		}
		matchInfoIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
