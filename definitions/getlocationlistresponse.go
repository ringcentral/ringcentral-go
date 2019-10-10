package definitions

// GetLocationListResponse Get Location List Response
type GetLocationListResponse struct {
	Records []LocationInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
