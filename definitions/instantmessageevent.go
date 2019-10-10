package definitions

// InstantMessageEvent Instant Message Event
type InstantMessageEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body InstantMessageEventBody `json:"body"`
}
