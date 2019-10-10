package definitions

// GreetingInfo Greeting Info
type GreetingInfo struct {
	Type string `json:"type"`
	UsageType string `json:"usageType"`
	Preset PresetInfo `json:"preset"`
}
