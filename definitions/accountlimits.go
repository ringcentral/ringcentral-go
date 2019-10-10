package definitions

// AccountLimits Account Limits
type AccountLimits struct {
	FreeSoftPhoneLinesPerExtension int `json:"freeSoftPhoneLinesPerExtension"`
	MeetingSize int `json:"meetingSize"`
	CloudRecordingStorage int `json:"cloudRecordingStorage"`
	MaxMonitoredExtensionsPerUser int `json:"maxMonitoredExtensionsPerUser"`
	MaxExtensionNumberLength int `json:"maxExtensionNumberLength"`
	SiteCodeLength int `json:"siteCodeLength"`
	ShortExtensionNumberLength int `json:"shortExtensionNumberLength"`
}
