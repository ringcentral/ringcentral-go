package definitions

// ForwardingInfo Forwarding Info
type ForwardingInfo struct {
	NotifyMySoftPhones bool `json:"notifyMySoftPhones"`
	NotifyAdminSoftPhones bool `json:"notifyAdminSoftPhones"`
	SoftPhonesRingCount int `json:"softPhonesRingCount"`
	RingingMode string `json:"ringingMode"`
	Rules []RuleInfo `json:"rules"`
	MobileTimeout bool `json:"mobileTimeout"`
}
