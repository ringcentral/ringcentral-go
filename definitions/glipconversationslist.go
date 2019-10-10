package definitions

// GlipConversationsList Glip Conversations List
type GlipConversationsList struct {
	Records []GlipConversationInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
