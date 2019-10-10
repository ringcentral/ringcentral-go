package definitions

// ShippingAddressInfo Shipping Address Info
type ShippingAddressInfo struct {
	CustomerName string `json:"customerName"`
	Street string `json:"street"`
	Street2 string `json:"street2"`
	City string `json:"city"`
	State string `json:"state"`
	StateId string `json:"stateId"`
	StateIsoCode string `json:"stateIsoCode"`
	StateName string `json:"stateName"`
	CountryId string `json:"countryId"`
	CountryIsoCode string `json:"countryIsoCode"`
	Country string `json:"country"`
	CountryName string `json:"countryName"`
	Zip string `json:"zip"`
}
