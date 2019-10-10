package definitions

// CompanyAnsweringRuleRequest Company Answering Rule Request
type CompanyAnsweringRuleRequest struct {
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
	Type string `json:"type"`
	Callers []CompanyAnsweringRuleCallersInfoRequest `json:"callers"`
	CalledNumbers []CompanyAnsweringRuleCalledNumberInfo `json:"calledNumbers"`
	Schedule CompanyAnsweringRuleScheduleInfoRequest `json:"schedule"`
	CallHandlingAction string `json:"callHandlingAction"`
	Extension CompanyAnsweringRuleExtensionInfo `json:"extension"`
	Greetings []GreetingInfo `json:"greetings"`
}
