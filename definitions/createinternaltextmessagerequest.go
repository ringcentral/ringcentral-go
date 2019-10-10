package definitions

// CreateInternalTextMessageRequest Create Internal Text Message Request
type CreateInternalTextMessageRequest struct {
	From PagerCallerInfoRequest `json:"from"`
	ReplyOn int `json:"replyOn"`
	Text string `json:"text"`
	To []PagerCallerInfoRequest `json:"to"`
}
