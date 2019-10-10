package definitions

// InboundTextsInfo Inbound Texts Info
type InboundTextsInfo struct {
	NotifyByEmail bool `json:"notifyByEmail"`
	NotifyBySms bool `json:"notifyBySms"`
	AdvancedEmailAddresses []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
}
