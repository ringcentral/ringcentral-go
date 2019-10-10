package definitions

// GlipChatInfo Glip Chat Info
type GlipChatInfo struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Public bool `json:"public"`
	Name string `json:"name"`
	Description string `json:"description"`
	Status string `json:"status"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Members []GlipChatMemberInfo `json:"members"`
}
