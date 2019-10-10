package definitions

// CreateSwitchInfo Create Switch Info
type CreateSwitchInfo struct {
	ChassisId string `json:"chassisId"`
	Name string `json:"name"`
	Site SwitchSiteInfo `json:"site"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
