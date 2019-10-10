package definitions

// CallMonitoringGroups Call Monitoring Groups
type CallMonitoringGroups struct {
	Uri string `json:"uri"`
	Records []CallMonitoringGroup `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
