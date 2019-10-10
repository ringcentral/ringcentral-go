package definitions

// DetailedExtensionPresenceEvent Detailed Extension Presence Event
type DetailedExtensionPresenceEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body DetailedExtensionPresenceEventBody `json:"body"`
}
