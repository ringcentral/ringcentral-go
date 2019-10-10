package definitions

// UpdateAnsweringRuleRequest Update Answering Rule Request
type UpdateAnsweringRuleRequest struct {
	Forwarding ForwardingInfoCreateRuleRequest `json:"forwarding"`
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Callers []CallersInfoRequest `json:"callers"`
	CalledNumbers []CalledNumberInfo `json:"calledNumbers"`
	Schedule ScheduleInfo `json:"schedule"`
	CallHandlingAction string `json:"callHandlingAction"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding"`
	Queue QueueInfo `json:"queue"`
	Voicemail VoicemailInfo `json:"voicemail"`
	Greetings []GreetingInfo `json:"greetings"`
	Screening string `json:"screening"`
	ShowInactiveNumbers bool `json:"showInactiveNumbers"`
}
