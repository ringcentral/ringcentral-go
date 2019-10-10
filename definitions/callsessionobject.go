package definitions

// CallSessionObject Call Session Object
type CallSessionObject struct {
	Id string `json:"id"`
	Origin OriginInfo `json:"origin"`
	VoiceCallToken string `json:"voiceCallToken"`
	Parties []CallParty `json:"parties"`
	CreationTime string `json:"creationTime"`
}
