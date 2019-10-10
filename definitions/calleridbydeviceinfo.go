package definitions

// CallerIdByDeviceInfo Caller Id By Device Info
type CallerIdByDeviceInfo struct {
	Type string `json:"type"`
	PhoneInfo CallerIdPhoneInfo `json:"phoneInfo"`
}
