package definitions

// MessageBody Message Body
type MessageBody struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	Attachments []MessageAttachmentInfo `json:"attachments"`
	Availability string `json:"availability"`
	ConversationId int `json:"conversationId"`
	Conversation ConversationInfo `json:"conversation"`
	CreationTime string `json:"creationTime"`
	DeliveryErrorCode string `json:"deliveryErrorCode"`
	Direction string `json:"direction"`
	FaxPageCount int `json:"faxPageCount"`
	FaxResolution string `json:"faxResolution"`
	From MessageSenderInfo `json:"from"`
	LastModifiedTime string `json:"lastModifiedTime"`
	MessageStatus string `json:"messageStatus"`
	PgToDepartment bool `json:"pgToDepartment"`
	Priority string `json:"priority"`
	ReadStatus string `json:"readStatus"`
	SmsDeliveryTime string `json:"smsDeliveryTime"`
	SmsSendingAttemptsCount int `json:"smsSendingAttemptsCount"`
	Subject string `json:"subject"`
	To []MessageRecipientInfo `json:"to"`
	Type string `json:"type"`
	VmTranscriptionStatus string `json:"vmTranscriptionStatus"`
}
