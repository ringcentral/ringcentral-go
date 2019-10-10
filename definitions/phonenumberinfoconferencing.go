package definitions

// PhoneNumberInfoConferencing Phone Number Info Conferencing
type PhoneNumberInfoConferencing struct {
	Country GetCountryInfoConferencing `json:"country"`
	Default bool `json:"default"`
	HasGreeting bool `json:"hasGreeting"`
	Location string `json:"location"`
	PhoneNumber string `json:"phoneNumber"`
	Premium bool `json:"premium"`
}
