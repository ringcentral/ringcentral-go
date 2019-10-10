package definitions

// DialInNumberResource Dial In Number Resource
type DialInNumberResource struct {
	PhoneNumber string `json:"phoneNumber"`
	FormattedNumber string `json:"formattedNumber"`
	Location string `json:"location"`
	Country CountryResource `json:"country"`
}
