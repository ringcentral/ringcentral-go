package definitions

// GetCountryInfoConferencing Get Country Info Conferencing
type GetCountryInfoConferencing struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	CallingCode string `json:"callingCode"`
	EmergencyCalling bool `json:"emergencyCalling"`
	IsoCode string `json:"isoCode"`
	Name string `json:"name"`
}
