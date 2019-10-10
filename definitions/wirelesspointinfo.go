package definitions

// WirelessPointInfo Wireless Point Info
type WirelessPointInfo struct {
	Id string `json:"id"`
	Bssid string `json:"bssid"`
	Name string `json:"name"`
	Site AutomaticLocationUpdatesSiteInfo `json:"site"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfo `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
