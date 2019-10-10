package definitions

// GetPresenceInfo Get Presence Info
type GetPresenceInfo struct {
	Uri string `json:"uri"`
	AllowSeeMyPresence bool `json:"allowSeeMyPresence"`
	DndStatus string `json:"dndStatus"`
	Extension GetPresenceExtensionInfo `json:"extension"`
	Message string `json:"message"`
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold"`
	PresenceStatus string `json:"presenceStatus"`
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall"`
	TelephonyStatus string `json:"telephonyStatus"`
	UserStatus string `json:"userStatus"`
	MeetingStatus string `json:"meetingStatus"`
	ActiveCalls []ActiveCallInfo `json:"activeCalls"`
}
