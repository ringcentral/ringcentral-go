package definitions

// RuleInfo Rule Info
type RuleInfo struct {
	Index int `json:"index"`
	RingCount int `json:"ringCount"`
	Enabled bool `json:"enabled"`
	ForwardingNumbers []CreateAnsweringRuleForwardingNumberInfo `json:"forwardingNumbers"`
}
