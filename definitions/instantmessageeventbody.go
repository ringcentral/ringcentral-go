package definitions

// InstantMessageEventBody Instant Message Event Body
type InstantMessageEventBody struct {
	Id string `json:"id"`
	To []NotificationRecipientInfo `json:"to"`
	From SenderInfo `json:"from"`
	Type string `json:"type"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	ReadStatus string `json:"readStatus"`
	Priority string `json:"priority"`
	Attachments []MessageAttachmentInfo `json:"attachments"`
	Direction string `json:"direction"`
	Availability string `json:"availability"`
	Subject string `json:"subject"`
	MessageStatus string `json:"messageStatus"`
	ConversationId string `json:"conversationId"`
	Conversation ConversationInfo `json:"conversation"`
	OwnerId string `json:"ownerId"`
}
