package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"match/x/match/types"
)

func TestGenesisState_Validate(t *testing.T) {
	tests := []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				MatchInfoList: []types.MatchInfo{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				MatchInfoCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated matchInfo",
			genState: &types.GenesisState{
				MatchInfoList: []types.MatchInfo{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid matchInfo count",
			genState: &types.GenesisState{
				MatchInfoList: []types.MatchInfo{
					{
						Id: 1,
					},
				},
				MatchInfoCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
