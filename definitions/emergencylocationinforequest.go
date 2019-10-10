package definitions

// EmergencyLocationInfoRequest Emergency Location Info Request
type EmergencyLocationInfoRequest struct {
	Id string `json:"id"`
	Address EmergencyLocationAddressInfoRequest `json:"address"`
}
