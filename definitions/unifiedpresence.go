package definitions

// UnifiedPresence Unified Presence
type UnifiedPresence struct {
	Status string `json:"status"`
	Glip UnifiedPresenceGlip `json:"glip"`
	Telephony UnifiedPresenceTelephony `json:"telephony"`
	Meeting UnifiedPresenceMeeting `json:"meeting"`
}
