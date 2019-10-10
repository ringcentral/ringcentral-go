package definitions

// CallerIdByFeatureInfo Caller Id By Feature Info
type CallerIdByFeatureInfo struct {
	Type string `json:"type"`
	PhoneInfo CallerIdPhoneInfo `json:"phoneInfo"`
}
