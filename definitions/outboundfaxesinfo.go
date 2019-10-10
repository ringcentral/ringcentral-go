package definitions

// OutboundFaxesInfo Outbound Faxes Info
type OutboundFaxesInfo struct {
	NotifyByEmail bool `json:"notifyByEmail"`
	NotifyBySms bool `json:"notifyBySms"`
	AdvancedEmailAddresses []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
}
