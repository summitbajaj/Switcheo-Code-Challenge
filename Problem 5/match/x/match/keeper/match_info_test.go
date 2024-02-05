package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "match/testutil/keeper"
	"match/testutil/nullify"
	"match/x/match/keeper"
	"match/x/match/types"
)

func createNMatchInfo(keeper keeper.Keeper, ctx context.Context, n int) []types.MatchInfo {
	items := make([]types.MatchInfo, n)
	for i := range items {
		items[i].Id = keeper.AppendMatchInfo(ctx, items[i])
	}
	return items
}

func TestMatchInfoGet(t *testing.T) {
	keeper, ctx := keepertest.MatchKeeper(t)
	items := createNMatchInfo(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetMatchInfo(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestMatchInfoRemove(t *testing.T) {
	keeper, ctx := keepertest.MatchKeeper(t)
	items := createNMatchInfo(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveMatchInfo(ctx, item.Id)
		_, found := keeper.GetMatchInfo(ctx, item.Id)
		require.False(t, found)
	}
}

func TestMatchInfoGetAll(t *testing.T) {
	keeper, ctx := keepertest.MatchKeeper(t)
	items := createNMatchInfo(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllMatchInfo(ctx)),
	)
}

func TestMatchInfoCount(t *testing.T) {
	keeper, ctx := keepertest.MatchKeeper(t)
	items := createNMatchInfo(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetMatchInfoCount(ctx))
}
