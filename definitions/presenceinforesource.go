package definitions

// PresenceInfoResource Presence Info Resource
type PresenceInfoResource struct {
	UserStatus string `json:"userStatus"`
	DndStatus string `json:"dndStatus"`
	Message string `json:"message"`
	AllowSeeMyPresence bool `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold"`
	ActiveCalls []ActiveCallInfo `json:"activeCalls"`
}
