package definitions

// CountryResource Country Resource
type CountryResource struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	Name string `json:"name"`
	IsoCode string `json:"isoCode"`
	CallingCode string `json:"callingCode"`
	EmergencyCalling bool `json:"emergencyCalling"`
	NumberSelling bool `json:"numberSelling"`
	LoginAllowed bool `json:"loginAllowed"`
}
