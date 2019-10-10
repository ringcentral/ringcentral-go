package definitions

// WirelessPointsList Wireless Points List
type WirelessPointsList struct {
	Uri string `json:"uri"`
	Records []WirelessPointInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
