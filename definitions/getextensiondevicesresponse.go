package definitions

// GetExtensionDevicesResponse Get Extension Devices Response
type GetExtensionDevicesResponse struct {
	Records []ExtensionDeviceResponse `json:"records"`
	Navigation DeviceProvisioningNavigationInfo `json:"navigation"`
	Paging DeviceProvisioningPagingInfo `json:"paging"`
}
