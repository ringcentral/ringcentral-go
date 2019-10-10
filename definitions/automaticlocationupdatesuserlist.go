package definitions

// AutomaticLocationUpdatesUserList Automatic Location Updates User List
type AutomaticLocationUpdatesUserList struct {
	Uri string `json:"uri"`
	Records []AutomaticLocationUpdatesUserInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
