package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"match/x/match/types"
)

func (k Keeper) MatchInfoAll(ctx context.Context, req *types.QueryAllMatchInfoRequest) (*types.QueryAllMatchInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var matchInfos []types.MatchInfo

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	matchInfoStore := prefix.NewStore(store, types.KeyPrefix(types.MatchInfoKey))

	pageRes, err := query.Paginate(matchInfoStore, req.Pagination, func(key []byte, value []byte) error {
		var matchInfo types.MatchInfo
		if err := k.cdc.Unmarshal(value, &matchInfo); err != nil {
			return err
		}

		matchInfos = append(matchInfos, matchInfo)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllMatchInfoResponse{MatchInfo: matchInfos, Pagination: pageRes}, nil
}

func (k Keeper) MatchInfo(ctx context.Context, req *types.QueryGetMatchInfoRequest) (*types.QueryGetMatchInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	matchInfo, found := k.GetMatchInfo(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetMatchInfoResponse{MatchInfo: matchInfo}, nil
}
