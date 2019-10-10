package definitions

// GetStateListResponse Get State List Response
type GetStateListResponse struct {
	Records []GetStateInfoResponse `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
