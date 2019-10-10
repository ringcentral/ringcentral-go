package definitions

// AutomaticRecordingResource Automatic Recording Resource
type AutomaticRecordingResource struct {
	Enabled bool `json:"enabled"`
	OutboundCallTones bool `json:"outboundCallTones"`
	OutboundCallAnnouncement bool `json:"outboundCallAnnouncement"`
	AllowMute bool `json:"allowMute"`
	ExtensionCount int `json:"extensionCount"`
}
