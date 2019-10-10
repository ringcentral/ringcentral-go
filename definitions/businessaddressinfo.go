package definitions

// BusinessAddressInfo Business Address Info
type BusinessAddressInfo struct {
	Country string `json:"country"`
	State string `json:"state"`
	City string `json:"city"`
	Street string `json:"street"`
	Zip string `json:"zip"`
}
