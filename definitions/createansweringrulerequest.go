package definitions

// CreateAnsweringRuleRequest Create Answering Rule Request
type CreateAnsweringRuleRequest struct {
	Enabled bool `json:"enabled"`
	Type string `json:"type"`
	Name string `json:"name"`
	Callers []CallersInfoRequest `json:"callers"`
	CalledNumbers []CalledNumberInfo `json:"calledNumbers"`
	Schedule ScheduleInfo `json:"schedule"`
	CallHandlingAction string `json:"callHandlingAction"`
	Forwarding ForwardingInfo `json:"forwarding"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding"`
	Queue QueueInfo `json:"queue"`
	Transfer TransferredExtensionInfo `json:"transfer"`
	Voicemail VoicemailInfo `json:"voicemail"`
	Greetings []GreetingInfo `json:"greetings"`
	Screening string `json:"screening"`
}
