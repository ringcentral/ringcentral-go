package definitions

// InboundMessageEvent Inbound Message Event
type InboundMessageEvent struct {
	Aps NotificationInfo `json:"aps"`
	MessageId string `json:"messageId"`
	ConversationId string `json:"conversationId"`
	From string `json:"from"`
	To string `json:"to"`
	OwnerId string `json:"ownerId"`
}
