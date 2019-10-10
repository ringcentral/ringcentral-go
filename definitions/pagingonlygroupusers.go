package definitions

// PagingOnlyGroupUsers Paging Only Group Users
type PagingOnlyGroupUsers struct {
	Records []PagingGroupExtensionInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
