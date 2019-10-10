package definitions

// ExtensionInfoEventBody Extension Info Event Body
type ExtensionInfoEventBody struct {
	ExtensionId string `json:"extensionId"`
	EventType string `json:"eventType"`
	Hints []string `json:"hints"`
	OwnerId string `json:"ownerId"`
}
