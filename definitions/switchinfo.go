package definitions

// SwitchInfo Switch Info
type SwitchInfo struct {
	Id string `json:"id"`
	ChassisId string `json:"chassisId"`
	Name string `json:"name"`
	Site SwitchSiteInfo `json:"site"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfo `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
