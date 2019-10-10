package definitions

// GetCountryListResponse Get Country List Response
type GetCountryListResponse struct {
	Records []GetCountryInfoDictionaryResponse `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
