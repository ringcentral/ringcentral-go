package definitions

// ForwardingInfoCreateRuleRequest Forwarding Info Create Rule Request
type ForwardingInfoCreateRuleRequest struct {
	NotifyMySoftPhones bool `json:"notifyMySoftPhones"`
	NotifyAdminSoftPhones bool `json:"notifyAdminSoftPhones"`
	SoftPhonesRingCount int `json:"softPhonesRingCount"`
	RingingMode string `json:"ringingMode"`
	Rules []RuleInfoCreateRuleRequest `json:"rules"`
	MobileTimeout bool `json:"mobileTimeout"`
}
