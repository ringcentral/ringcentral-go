package definitions

// CompanyPhoneNumberInfo Company Phone Number Info
type CompanyPhoneNumberInfo struct {
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
	TemporaryNumber TemporaryNumberInfo `json:"temporaryNumber"`
	ContactCenterProvider ContactCenterProvider `json:"contactCenterProvider"`
}
