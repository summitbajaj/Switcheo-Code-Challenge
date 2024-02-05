package keeper_test

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "match/testutil/keeper"
	"match/testutil/nullify"
	"match/x/match/types"
)

func TestMatchInfoQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.MatchKeeper(t)
	msgs := createNMatchInfo(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetMatchInfoRequest
		response *types.QueryGetMatchInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetMatchInfoRequest{Id: msgs[0].Id},
			response: &types.QueryGetMatchInfoResponse{MatchInfo: msgs[0]},
		},
		{
			desc:     "Second",
			request:  &types.QueryGetMatchInfoRequest{Id: msgs[1].Id},
			response: &types.QueryGetMatchInfoResponse{MatchInfo: msgs[1]},
		},
		{
			desc:    "KeyNotFound",
			request: &types.QueryGetMatchInfoRequest{Id: uint64(len(msgs))},
			err:     sdkerrors.ErrKeyNotFound,
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.MatchInfo(ctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestMatchInfoQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.MatchKeeper(t)
	msgs := createNMatchInfo(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllMatchInfoRequest {
		return &types.QueryAllMatchInfoRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MatchInfoAll(ctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MatchInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MatchInfo),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.MatchInfoAll(ctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.MatchInfo), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.MatchInfo),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.MatchInfoAll(ctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.MatchInfo),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.MatchInfoAll(ctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
