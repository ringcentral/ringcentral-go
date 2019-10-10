package definitions

// UnifiedPresenceListEntry Unified Presence List Entry
type UnifiedPresenceListEntry struct {
	ResourceId string `json:"resourceId"`
	Status int `json:"status"`
	Body UnifiedPresence `json:"body"`
}
