package definitions

// EmergencyLocationList Emergency Location List
type EmergencyLocationList struct {
	Records []EmergencyLocationInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
