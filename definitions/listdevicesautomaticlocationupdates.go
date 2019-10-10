package definitions

// ListDevicesAutomaticLocationUpdates List Devices Automatic Location Updates
type ListDevicesAutomaticLocationUpdates struct {
	Uri string `json:"uri"`
	Records []AutomaticLocationUpdatesDeviceInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
