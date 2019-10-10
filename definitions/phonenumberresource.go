package definitions

// PhoneNumberResource Phone Number Resource
type PhoneNumberResource struct {
	FormattedPhoneNumber string `json:"formattedPhoneNumber"`
	PhoneNumber string `json:"phoneNumber"`
	Type string `json:"type"`
	Label string `json:"label"`
	UsageType string `json:"usageType"`
}
