package definitions

// GetCountryInfoNumberParser Get Country Info Number Parser
type GetCountryInfoNumberParser struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	CallingCode string `json:"callingCode"`
	EmergencyCalling bool `json:"emergencyCalling"`
	IsoCode string `json:"isoCode"`
	Name string `json:"name"`
}
