package definitions

// AutomaticLocationUpdatesDeviceInfo Automatic Location Updates Device Info
type AutomaticLocationUpdatesDeviceInfo struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Serial string `json:"serial"`
	FeatureEnabled bool `json:"featureEnabled"`
	Name string `json:"name"`
	Model AutomaticLocationUpdatesModelInfo `json:"model"`
	Site string `json:"site"`
	PhoneLines []AutomaticLocationUpdatesPhoneLine `json:"phoneLines"`
}
