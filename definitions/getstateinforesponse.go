package definitions

// GetStateInfoResponse Get State Info Response
type GetStateInfoResponse struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Country GetCountryInfoState `json:"country"`
	IsoCode string `json:"isoCode"`
	Name string `json:"name"`
}
