package match_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "match/testutil/keeper"
	"match/testutil/nullify"
	"match/x/match/module"
	"match/x/match/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MatchKeeper(t)
	match.InitGenesis(ctx, k, genesisState)
	got := match.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
