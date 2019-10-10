package definitions

// AccountPresenceEvent Account Presence Event
type AccountPresenceEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body AccountPresenceEventBody `json:"body"`
}
