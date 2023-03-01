package keeper

import (
	"encoding/binary"

	"exabits/x/computing/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetDevicesCount get the total number of devices
func (k Keeper) GetDevicesCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DevicesCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetDevicesCount set the total number of devices
func (k Keeper) SetDevicesCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.DevicesCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendDevices appends a devices in the store with a new id and update the count
func (k Keeper) AppendDevices(
	ctx sdk.Context,
	devices types.Devices,
) uint64 {
	// Create the devices
	count := k.GetDevicesCount(ctx)

	// Set the ID of the appended value
	devices.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DevicesKey))
	appendedValue := k.cdc.MustMarshal(&devices)
	store.Set(GetDevicesIDBytes(devices.Id), appendedValue)

	// Update devices count
	k.SetDevicesCount(ctx, count+1)

	return count
}

// SetDevices set a specific devices in the store
func (k Keeper) SetDevices(ctx sdk.Context, devices types.Devices) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DevicesKey))
	b := k.cdc.MustMarshal(&devices)
	store.Set(GetDevicesIDBytes(devices.Id), b)
}

// GetDevices returns a devices from its id
func (k Keeper) GetDevices(ctx sdk.Context, id uint64) (val types.Devices, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DevicesKey))
	b := store.Get(GetDevicesIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDevices removes a devices from the store
func (k Keeper) RemoveDevices(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DevicesKey))
	store.Delete(GetDevicesIDBytes(id))
}

// GetAllDevices returns all devices
func (k Keeper) GetAllDevices(ctx sdk.Context) (list []types.Devices) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DevicesKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Devices
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetDevicesIDBytes returns the byte representation of the ID
func GetDevicesIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetDevicesIDFromBytes returns ID in uint64 format from a byte array
func GetDevicesIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
