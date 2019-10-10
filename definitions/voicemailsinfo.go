package definitions

// VoicemailsInfo Voicemails Info
type VoicemailsInfo struct {
	NotifyByEmail bool `json:"notifyByEmail"`
	NotifyBySms bool `json:"notifyBySms"`
	AdvancedEmailAddresses []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
	IncludeAttachment bool `json:"includeAttachment"`
	MarkAsRead bool `json:"markAsRead"`
}
