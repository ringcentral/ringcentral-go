package definitions

// UserPhoneNumberInfo User Phone Number Info
type UserPhoneNumberInfo struct {
	Id string `json:"id"`
	Country CountryInfo `json:"country"`
	Extension ExtensionInfo `json:"extension"`
	Label string `json:"label"`
	Location string `json:"location"`
	PaymentType string `json:"paymentType"`
	PhoneNumber string `json:"phoneNumber"`
	Status string `json:"status"`
	Type string `json:"type"`
	UsageType string `json:"usageType"`
	Features []string `json:"features"`
}
