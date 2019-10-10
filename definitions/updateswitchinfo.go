package definitions

// UpdateSwitchInfo Update Switch Info
type UpdateSwitchInfo struct {
	Id string `json:"id"`
	ChassisId string `json:"chassisId"`
	Name string `json:"name"`
	Site SwitchSiteInfo `json:"site"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
