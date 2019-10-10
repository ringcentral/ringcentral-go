package definitions

// UpdateUnifiedPresence Update Unified Presence
type UpdateUnifiedPresence struct {
	Glip UpdateUnifiedPresenceGlip `json:"glip"`
	Telephony UpdateUnifiedPresenceTelephony `json:"telephony"`
}
