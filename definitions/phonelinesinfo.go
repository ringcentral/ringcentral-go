package definitions

// PhoneLinesInfo Phone Lines Info
type PhoneLinesInfo struct {
	LineType string `json:"lineType"`
	PhoneInfo PhoneNumberInfoIntId `json:"phoneInfo"`
}
