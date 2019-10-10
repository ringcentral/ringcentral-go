package definitions

// NetworksList Networks List
type NetworksList struct {
	Uri string `json:"uri"`
	Records []NetworkInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
