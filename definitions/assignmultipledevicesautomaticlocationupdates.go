package definitions

// AssignMultipleDevicesAutomaticLocationUpdates Assign Multiple Devices Automatic Location Updates
type AssignMultipleDevicesAutomaticLocationUpdates struct {
	EnabledDeviceIds []string `json:"enabledDeviceIds"`
	DisabledDeviceIds []string `json:"disabledDeviceIds"`
}
