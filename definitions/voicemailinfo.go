package definitions

// VoicemailInfo Voicemail Info
type VoicemailInfo struct {
	Enabled bool `json:"enabled"`
	Recipient RecipientInfo `json:"recipient"`
}
