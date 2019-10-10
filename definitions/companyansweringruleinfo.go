package definitions

// CompanyAnsweringRuleInfo Company Answering Rule Info
type CompanyAnsweringRuleInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Enabled bool `json:"enabled"`
	Type string `json:"type"`
	Name string `json:"name"`
	Callers []CompanyAnsweringRuleCallersInfoRequest `json:"callers"`
	CalledNumbers []CompanyAnsweringRuleCalledNumberInfoRequest `json:"calledNumbers"`
	Schedule CompanyAnsweringRuleScheduleInfo `json:"schedule"`
	CallHandlingAction string `json:"callHandlingAction"`
	Extension CompanyAnsweringRuleCallersInfoRequest `json:"extension"`
	Greetings []GreetingInfo `json:"greetings"`
}
