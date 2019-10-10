package definitions

// PhoneNumberInfoIntId Phone Number Info Int Id
type PhoneNumberInfoIntId struct {
	Id int `json:"id"`
	Country PhoneNumberCountryInfo `json:"country"`
	Extension ExtensionInfo `json:"extension"`
	Label string `json:"label"`
	Location string `json:"location"`
	PaymentType string `json:"paymentType"`
	PhoneNumber string `json:"phoneNumber"`
	Status string `json:"status"`
	Type string `json:"type"`
	UsageType string `json:"usageType"`
}
