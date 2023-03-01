package keeper

import (
	"context"
	"fmt"

	"exabits/x/computing/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDevices(goCtx context.Context, msg *types.MsgCreateDevices) (*types.MsgCreateDevicesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var devices = types.Devices{
		Creator:      msg.Creator,
		Address:      msg.Address,
		MachineId:    msg.MachineId,
		MachineToken: msg.MachineToken,
		Coords:       msg.Coords,
		Ip:           msg.Ip,
	}

	id := k.AppendDevices(
		ctx,
		devices,
	)

	return &types.MsgCreateDevicesResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateDevices(goCtx context.Context, msg *types.MsgUpdateDevices) (*types.MsgUpdateDevicesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var devices = types.Devices{
		Creator:      msg.Creator,
		Id:           msg.Id,
		Address:      msg.Address,
		MachineId:    msg.MachineId,
		MachineToken: msg.MachineToken,
		Coords:       msg.Coords,
		Ip:           msg.Ip,
	}

	// Checks that the element exists
	val, found := k.GetDevices(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetDevices(ctx, devices)

	return &types.MsgUpdateDevicesResponse{}, nil
}

func (k msgServer) DeleteDevices(goCtx context.Context, msg *types.MsgDeleteDevices) (*types.MsgDeleteDevicesResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetDevices(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveDevices(ctx, msg.Id)

	return &types.MsgDeleteDevicesResponse{}, nil
}
