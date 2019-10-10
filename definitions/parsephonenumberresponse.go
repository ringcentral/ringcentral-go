package definitions

// ParsePhoneNumberResponse Parse Phone Number Response
type ParsePhoneNumberResponse struct {
	Uri string `json:"uri"`
	HomeCountry GetCountryInfoNumberParser `json:"homeCountry"`
	PhoneNumbers []PhoneNumberInfoNumberParser `json:"phoneNumbers"`
}
