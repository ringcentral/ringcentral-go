package definitions

// AccountPresenceEventBody Account Presence Event Body
type AccountPresenceEventBody struct {
	ExtensionId string `json:"extensionId"`
	TelephonyStatus string `json:"telephonyStatus"`
	Sequence int `json:"sequence"`
	PresenceStatus string `json:"presenceStatus"`
	UserStatus string `json:"userStatus"`
	DndStatus string `json:"dndStatus"`
	AllowSeeMyPresence bool `json:"allowSeeMyPresence"`
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall"`
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold"`
	OwnerId string `json:"ownerId"`
}
