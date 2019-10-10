package definitions

// DeviceProvisioningNavigationInfo Device Provisioning Navigation Info
type DeviceProvisioningNavigationInfo struct {
	FirstPage DeviceProvisioningNavigationInfoUri `json:"firstPage"`
	NextPage DeviceProvisioningNavigationInfoUri `json:"nextPage"`
	PreviousPage DeviceProvisioningNavigationInfoUri `json:"previousPage"`
	LastPage DeviceProvisioningNavigationInfoUri `json:"lastPage"`
}
