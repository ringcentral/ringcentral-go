package definitions

// ExtensionListEvent Extension List Event
type ExtensionListEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body ExtensionListEventBody `json:"body"`
}
