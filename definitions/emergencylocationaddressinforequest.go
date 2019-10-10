package definitions

// EmergencyLocationAddressInfoRequest Emergency Location Address Info Request
type EmergencyLocationAddressInfoRequest struct {
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Street string `json:"street"`
	Street2 string `json:"street2"`
	Zip string `json:"zip"`
}
