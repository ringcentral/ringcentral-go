package definitions

// GetTimezoneListResponse Get Timezone List Response
type GetTimezoneListResponse struct {
	Records []GetTimezoneInfoResponse `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
