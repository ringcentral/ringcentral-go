package definitions

// AccountDeviceUpdate Account Device Update
type AccountDeviceUpdate struct {
	EmergencyServiceAddress EmergencyServiceAddressResourceRequest `json:"emergencyServiceAddress"`
	Extension DeviceUpdateExtensionInfo `json:"extension"`
	PhoneLines DeviceUpdatePhoneLinesInfo `json:"phoneLines"`
	UseAsCommonPhone bool `json:"useAsCommonPhone"`
}
