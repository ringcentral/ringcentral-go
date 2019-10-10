package definitions

// CallMonitoringGroupMemberList Call Monitoring Group Member List
type CallMonitoringGroupMemberList struct {
	Uri string `json:"uri"`
	Records []CallMonitoringGroupMemberInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
