package definitions

// AutomaticLocationUpdatesPhoneLine Automatic Location Updates Phone Line
type AutomaticLocationUpdatesPhoneLine struct {
	LineType string `json:"lineType"`
	PhoneInfo AutomaticLocationUpdatesPhoneNumberInfo `json:"phoneInfo"`
}
