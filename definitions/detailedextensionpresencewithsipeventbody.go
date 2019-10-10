package definitions

// DetailedExtensionPresenceWithSIPEventBody Detailed Extension Presence With SIPEvent Body
type DetailedExtensionPresenceWithSIPEventBody struct {
	ExtensionId string `json:"extensionId"`
	TelephonyStatus string `json:"telephonyStatus"`
	ActiveCalls []ActiveCallInfo `json:"activeCalls"`
	Sequence int `json:"sequence"`
	PresenceStatus string `json:"presenceStatus"`
	UserStatus string `json:"userStatus"`
	DndStatus string `json:"dndStatus"`
	AllowSeeMyPresence bool `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold"`
	TotalActiveCalls int `json:"totalActiveCalls"`
	OwnerId string `json:"ownerId"`
}
