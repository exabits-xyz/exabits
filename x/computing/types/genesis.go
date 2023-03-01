package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		DevicesList: []Devices{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated ID in devices
	devicesIdMap := make(map[uint64]bool)
	devicesCount := gs.GetDevicesCount()
	for _, elem := range gs.DevicesList {
		if _, ok := devicesIdMap[elem.Id]; ok {
			return fmt.Errorf("duplicated id for devices")
		}
		if elem.Id >= devicesCount {
			return fmt.Errorf("devices id should be lower or equal than the last id")
		}
		devicesIdMap[elem.Id] = true
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
