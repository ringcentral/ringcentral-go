package definitions

// EmergencyServiceAddressResourceRequest Emergency Service Address Resource Request
type EmergencyServiceAddressResourceRequest struct {
	Street string `json:"street"`
	Street2 string `json:"street2"`
	City string `json:"city"`
	Zip string `json:"zip"`
	CustomerName string `json:"customerName"`
	State string `json:"state"`
	Country string `json:"country"`
}
