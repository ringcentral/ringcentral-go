package definitions

// RuleInfoCreateRuleRequest Rule Info Create Rule Request
type RuleInfoCreateRuleRequest struct {
	Index int `json:"index"`
	RingCount int `json:"ringCount"`
	Enabled bool `json:"enabled"`
	ForwardingNumbers []ForwardingNumberInfoRulesCreateRuleRequest `json:"forwardingNumbers"`
}
