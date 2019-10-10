package definitions

// MessageEventBody Message Event Body
type MessageEventBody struct {
	ExtensionId string `json:"extensionId"`
	LastUpdated string `json:"lastUpdated"`
	Changes []MessageChanges `json:"changes"`
	OwnerId string `json:"ownerId"`
}
