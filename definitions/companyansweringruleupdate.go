package definitions

// CompanyAnsweringRuleUpdate Company Answering Rule Update
type CompanyAnsweringRuleUpdate struct {
	Enabled bool `json:"enabled"`
	Name string `json:"name"`
	Callers []CompanyAnsweringRuleCallersInfoRequest `json:"callers"`
	CalledNumbers []CompanyAnsweringRuleCalledNumberInfo `json:"calledNumbers"`
	Schedule CompanyAnsweringRuleScheduleInfoRequest `json:"schedule"`
	CallHandlingAction string `json:"callHandlingAction"`
	Extension CompanyAnsweringRuleCallersInfoRequest `json:"extension"`
	Greetings []GreetingInfo `json:"greetings"`
}
