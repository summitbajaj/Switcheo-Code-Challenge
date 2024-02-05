package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"match/x/match/types"
)

func (k msgServer) CreateMatchInfo(goCtx context.Context, msg *types.MsgCreateMatchInfo) (*types.MsgCreateMatchInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var matchInfo = types.MatchInfo{
		Creator:   msg.Creator,
		Matchdate: msg.Matchdate,
		Home:      msg.Home,
		Away:      msg.Away,
	}

	id := k.AppendMatchInfo(
		ctx,
		matchInfo,
	)

	return &types.MsgCreateMatchInfoResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateMatchInfo(goCtx context.Context, msg *types.MsgUpdateMatchInfo) (*types.MsgUpdateMatchInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var matchInfo = types.MatchInfo{
		Creator:   msg.Creator,
		Id:        msg.Id,
		Matchdate: msg.Matchdate,
		Home:      msg.Home,
		Away:      msg.Away,
	}

	// Checks that the element exists
	val, found := k.GetMatchInfo(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetMatchInfo(ctx, matchInfo)

	return &types.MsgUpdateMatchInfoResponse{}, nil
}

func (k msgServer) DeleteMatchInfo(goCtx context.Context, msg *types.MsgDeleteMatchInfo) (*types.MsgDeleteMatchInfoResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetMatchInfo(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveMatchInfo(ctx, msg.Id)

	return &types.MsgDeleteMatchInfoResponse{}, nil
}
