package definitions

// WirelessPointInfoRequest Wireless Point Info Request
type WirelessPointInfoRequest struct {
	Id string `json:"id"`
	Bssid string `json:"bssid"`
	Name string `json:"name"`
	Site AutomaticLocationUpdatesSiteInfo `json:"site"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
