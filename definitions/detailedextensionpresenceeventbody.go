package definitions

// DetailedExtensionPresenceEventBody Detailed Extension Presence Event Body
type DetailedExtensionPresenceEventBody struct {
	ExtensionId string `json:"extensionId"`
	TelephonyStatus string `json:"telephonyStatus"`
	ActiveCalls []ActiveCallInfoWithoutSIP `json:"activeCalls"`
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
