package definitions

// EmergencyServiceAddressResource Emergency Service Address Resource
type EmergencyServiceAddressResource struct {
	Street string `json:"street"`
	Street2 string `json:"street2"`
	City string `json:"city"`
	Zip string `json:"zip"`
	CustomerName string `json:"customerName"`
	State string `json:"state"`
	StateId string `json:"stateId"`
	StateIsoCode string `json:"stateIsoCode"`
	StateName string `json:"stateName"`
	CountryId string `json:"countryId"`
	CountryIsoCode string `json:"countryIsoCode"`
	Country string `json:"country"`
	CountryName string `json:"countryName"`
	OutOfCountry bool `json:"outOfCountry"`
}
