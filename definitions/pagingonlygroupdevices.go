package definitions

// PagingOnlyGroupDevices Paging Only Group Devices
type PagingOnlyGroupDevices struct {
	Records []PagingDeviceInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
