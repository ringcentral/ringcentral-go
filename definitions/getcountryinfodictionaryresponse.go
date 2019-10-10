package definitions

// GetCountryInfoDictionaryResponse Get Country Info Dictionary Response
type GetCountryInfoDictionaryResponse struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	CallingCode string `json:"callingCode"`
	EmergencyCalling bool `json:"emergencyCalling"`
	IsoCode string `json:"isoCode"`
	Name string `json:"name"`
	NumberSelling bool `json:"numberSelling"`
	LoginAllowed bool `json:"loginAllowed"`
	SignupAllowed bool `json:"signupAllowed"`
	FreeSoftphoneLine bool `json:"freeSoftphoneLine"`
}
