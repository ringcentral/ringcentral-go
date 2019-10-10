package definitions

// GetExtensionListResponse Get Extension List Response
type GetExtensionListResponse struct {
	Records []GetExtensionInfoResponse `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
