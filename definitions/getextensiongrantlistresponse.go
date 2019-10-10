package definitions

// GetExtensionGrantListResponse Get Extension Grant List Response
type GetExtensionGrantListResponse struct {
	Records []GrantInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
