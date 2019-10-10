package definitions

// EmergencyLocationInfo Emergency Location Info
type EmergencyLocationInfo struct {
	Id string `json:"id"`
	Address EmergencyLocationAddressInfo `json:"address"`
}
