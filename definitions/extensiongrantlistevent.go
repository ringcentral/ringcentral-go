package definitions

// ExtensionGrantListEvent Extension Grant List Event
type ExtensionGrantListEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body ExtensionGrantListEventBody `json:"body"`
}
