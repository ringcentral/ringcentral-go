package definitions

// NotificationSettingsUpdateRequest Notification Settings Update Request
type NotificationSettingsUpdateRequest struct {
	EmailAddresses []string `json:"emailAddresses"`
	SmsEmailAddresses []string `json:"smsEmailAddresses"`
	AdvancedMode bool `json:"advancedMode"`
	Voicemails VoicemailsInfo `json:"voicemails"`
	InboundFaxes InboundFaxesInfo `json:"inboundFaxes"`
	OutboundFaxes OutboundFaxesInfo `json:"outboundFaxes"`
	InboundTexts InboundTextsInfo `json:"inboundTexts"`
	MissedCalls MissedCallsInfo `json:"missedCalls"`
}
