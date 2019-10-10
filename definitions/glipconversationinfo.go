package definitions

// GlipConversationInfo Glip Conversation Info
type GlipConversationInfo struct {
	Id string `json:"id"`
	Type string `json:"type"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Members []GlipChatMemberInfo `json:"members"`
}
