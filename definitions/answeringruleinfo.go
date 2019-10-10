package definitions

// AnsweringRuleInfo Answering Rule Info
type AnsweringRuleInfo struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
	Schedule ScheduleInfo `json:"schedule"`
	CalledNumbers []CalledNumberInfo `json:"calledNumbers"`
	Callers []CallersInfo `json:"callers"`
	CallHandlingAction string `json:"callHandlingAction"`
	Forwarding ForwardingInfo `json:"forwarding"`
	UnconditionalForwarding UnconditionalForwardingInfo `json:"unconditionalForwarding"`
	Queue QueueInfo `json:"queue"`
	Transfer TransferredExtensionInfo `json:"transfer"`
	Voicemail VoicemailInfo `json:"voicemail"`
	Greetings []GreetingInfo `json:"greetings"`
	Screening string `json:"screening"`
}
