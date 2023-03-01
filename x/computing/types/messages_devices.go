package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateDevices = "create_devices"
	TypeMsgUpdateDevices = "update_devices"
	TypeMsgDeleteDevices = "delete_devices"
)

var _ sdk.Msg = &MsgCreateDevices{}

func NewMsgCreateDevices(creator string, address string, machineId string, machineToken string, coords string, ip string) *MsgCreateDevices {
	return &MsgCreateDevices{
		Creator:      creator,
		Address:      address,
		MachineId:    machineId,
		MachineToken: machineToken,
		Coords:       coords,
		Ip:           ip,
	}
}

func (msg *MsgCreateDevices) Route() string {
	return RouterKey
}

func (msg *MsgCreateDevices) Type() string {
	return TypeMsgCreateDevices
}

func (msg *MsgCreateDevices) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDevices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDevices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDevices{}

func NewMsgUpdateDevices(creator string, id uint64, address string, machineId string, machineToken string, coords string, ip string) *MsgUpdateDevices {
	return &MsgUpdateDevices{
		Id:           id,
		Creator:      creator,
		Address:      address,
		MachineId:    machineId,
		MachineToken: machineToken,
		Coords:       coords,
		Ip:           ip,
	}
}

func (msg *MsgUpdateDevices) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDevices) Type() string {
	return TypeMsgUpdateDevices
}

func (msg *MsgUpdateDevices) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDevices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDevices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDevices{}

func NewMsgDeleteDevices(creator string, id uint64) *MsgDeleteDevices {
	return &MsgDeleteDevices{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteDevices) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDevices) Type() string {
	return TypeMsgDeleteDevices
}

func (msg *MsgDeleteDevices) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDevices) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDevices) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
