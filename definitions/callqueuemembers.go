package definitions

// CallQueueMembers Call Queue Members
type CallQueueMembers struct {
	Uri string `json:"uri"`
	Records []CallQueueMemberInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
