package definitions

// GlipChatsList Glip Chats List
type GlipChatsList struct {
	Records []GlipChatInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
