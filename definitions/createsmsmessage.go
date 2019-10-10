package definitions

// CreateSMSMessage Create SMSMessage
type CreateSMSMessage struct {
	From MessageStoreCallerInfoRequest `json:"from"`
	To []MessageStoreCallerInfoRequest `json:"to"`
	Text string `json:"text"`
}
