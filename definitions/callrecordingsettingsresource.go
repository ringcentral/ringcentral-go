package definitions

// CallRecordingSettingsResource Call Recording Settings Resource
type CallRecordingSettingsResource struct {
	OnDemand OnDemandResource `json:"onDemand"`
	Automatic AutomaticRecordingResource `json:"automatic"`
	Greetings []GreetingResource `json:"greetings"`
}
