package definitions

// PhoneLineResource Phone Line Resource
type PhoneLineResource struct {
	LineType string `json:"lineType"`
	PhoneInfo PhoneNumberResourceIntId `json:"phoneInfo"`
	EmergencyAddress EmergencyAddress `json:"emergencyAddress"`
}
