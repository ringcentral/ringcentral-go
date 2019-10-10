package definitions

// DevicePhoneNumberInfo Device Phone Number Info
type DevicePhoneNumberInfo struct {
	Id int `json:"id"`
	Country DevicePhoneNumberCountryInfo `json:"country"`
	PaymentType string `json:"paymentType"`
	PhoneNumber string `json:"phoneNumber"`
	UsageType string `json:"usageType"`
	Type string `json:"type"`
}
