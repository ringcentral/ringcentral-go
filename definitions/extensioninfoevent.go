package definitions

// ExtensionInfoEvent Extension Info Event
type ExtensionInfoEvent struct {
	Uuid string `json:"uuid"`
	Event string `json:"event"`
	Timestamp string `json:"timestamp"`
	SubscriptionId string `json:"subscriptionId"`
	Body ExtensionInfoEventBody `json:"body"`
}
