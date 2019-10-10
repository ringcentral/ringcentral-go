package definitions

// DetailedExtensionPresenceWithSIPEvent Detailed Extension Presence With SIPEvent
type DetailedExtensionPresenceWithSIPEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body DetailedExtensionPresenceWithSIPEventBody `json:"body"`
}
