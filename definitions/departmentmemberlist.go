package definitions

// DepartmentMemberList Department Member List
type DepartmentMemberList struct {
	Records []ExtensionInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
