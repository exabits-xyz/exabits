package keeper

import (
	"context"

	"exabits/x/computing/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) DevicesAll(goCtx context.Context, req *types.QueryAllDevicesRequest) (*types.QueryAllDevicesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var devicess []types.Devices
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	devicesStore := prefix.NewStore(store, types.KeyPrefix(types.DevicesKey))

	pageRes, err := query.Paginate(devicesStore, req.Pagination, func(key []byte, value []byte) error {
		var devices types.Devices
		if err := k.cdc.Unmarshal(value, &devices); err != nil {
			return err
		}

		devicess = append(devicess, devices)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllDevicesResponse{Devices: devicess, Pagination: pageRes}, nil
}

func (k Keeper) Devices(goCtx context.Context, req *types.QueryGetDevicesRequest) (*types.QueryGetDevicesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	devices, found := k.GetDevices(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetDevicesResponse{Devices: devices}, nil
}
