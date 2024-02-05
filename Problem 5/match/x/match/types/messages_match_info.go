package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateMatchInfo{}

func NewMsgCreateMatchInfo(creator string, matchdate string, home string, away string) *MsgCreateMatchInfo {
	return &MsgCreateMatchInfo{
		Creator:   creator,
		Matchdate: matchdate,
		Home:      home,
		Away:      away,
	}
}

func (msg *MsgCreateMatchInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateMatchInfo{}

func NewMsgUpdateMatchInfo(creator string, id uint64, matchdate string, home string, away string) *MsgUpdateMatchInfo {
	return &MsgUpdateMatchInfo{
		Id:        id,
		Creator:   creator,
		Matchdate: matchdate,
		Home:      home,
		Away:      away,
	}
}

func (msg *MsgUpdateMatchInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteMatchInfo{}

func NewMsgDeleteMatchInfo(creator string, id uint64) *MsgDeleteMatchInfo {
	return &MsgDeleteMatchInfo{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteMatchInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
