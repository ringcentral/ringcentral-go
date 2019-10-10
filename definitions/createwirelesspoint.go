package definitions

// CreateWirelessPoint Create Wireless Point
type CreateWirelessPoint struct {
	Bssid string `json:"bssid"`
	Name string `json:"name"`
	Site AutomaticLocationUpdatesSiteInfo `json:"site"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
