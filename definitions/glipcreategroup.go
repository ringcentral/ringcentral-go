package definitions

// GlipCreateGroup Glip Create Group
type GlipCreateGroup struct {
	Type string `json:"type"`
	IsPublic bool `json:"isPublic"`
	Name string `json:"name"`
	Description string `json:"description"`
	Members []string `json:"members"`
}
