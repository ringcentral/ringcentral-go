package definitions

// InboundFaxesInfo Inbound Faxes Info
type InboundFaxesInfo struct {
	NotifyByEmail bool `json:"notifyByEmail"`
	NotifyBySms bool `json:"notifyBySms"`
	AdvancedEmailAddresses []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
	IncludeAttachment bool `json:"includeAttachment"`
	MarkAsRead bool `json:"markAsRead"`
}
