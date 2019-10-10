package definitions

// ExtensionPresenceEvent Extension Presence Event
type ExtensionPresenceEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body ExtensionPresenceEventBody `json:"body"`
}
