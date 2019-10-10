package definitions

// DevicePhoneLinesInfo Device Phone Lines Info
type DevicePhoneLinesInfo struct {
	LineType string `json:"lineType"`
	EmergencyAddress DevicePhoneLinesEmergencyAddressInfo `json:"emergencyAddress"`
	PhoneInfo DevicePhoneNumberInfo `json:"phoneInfo"`
}
