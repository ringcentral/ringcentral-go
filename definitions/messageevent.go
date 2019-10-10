package definitions

// MessageEvent Message Event
type MessageEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body MessageEventBody `json:"body"`
}
