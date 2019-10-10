package definitions

// CallPartyReply Call Party Reply
type CallPartyReply struct {
	ReplyWithText string `json:"replyWithText"`
	ReplyWithPattern ReplyWithPattern `json:"replyWithPattern"`
}
