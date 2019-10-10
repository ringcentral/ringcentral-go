package definitions

// CallRecordingCustomGreeting Call Recording Custom Greeting
type CallRecordingCustomGreeting struct {
	Type string `json:"type"`
	Custom CallRecordingCustomGreetingData `json:"custom"`
	Language CallRecordingCustomGreetingLanguage `json:"language"`
}
