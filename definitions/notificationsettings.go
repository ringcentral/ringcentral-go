package definitions

// NotificationSettings Notification Settings
type NotificationSettings struct {
	Uri string `json:"uri"`
	EmailAddresses []string `json:"emailAddresses"`
	SmsEmailAddresses []string `json:"smsEmailAddresses"`
	AdvancedMode bool `json:"advancedMode"`
	Voicemails VoicemailsInfo `json:"voicemails"`
	InboundFaxes InboundFaxesInfo `json:"inboundFaxes"`
	OutboundFaxes OutboundFaxesInfo `json:"outboundFaxes"`
	InboundTexts InboundTextsInfo `json:"inboundTexts"`
	MissedCalls MissedCallsInfo `json:"missedCalls"`
}
