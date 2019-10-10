package definitions

// LocationUpdatesEmergencyAddressInfoRequest Location Updates Emergency Address Info Request
type LocationUpdatesEmergencyAddressInfoRequest struct {
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Street string `json:"street"`
	Street2 string `json:"street2"`
	Zip string `json:"zip"`
}
