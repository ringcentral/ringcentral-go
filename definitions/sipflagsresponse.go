package definitions

// SIPFlagsResponse SIPFlags Response
type SIPFlagsResponse struct {
	VoipFeatureEnabled string `json:"voipFeatureEnabled"`
	VoipCountryBlocked string `json:"voipCountryBlocked"`
	OutboundCallsEnabled string `json:"outboundCallsEnabled"`
	DscpEnabled bool `json:"dscpEnabled"`
	DscpSignaling int `json:"dscpSignaling"`
	DscpVoice int `json:"dscpVoice"`
	DscpVideo int `json:"dscpVideo"`
}
