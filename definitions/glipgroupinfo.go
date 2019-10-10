package definitions

// GlipGroupInfo Glip Group Info
type GlipGroupInfo struct {
	Id string `json:"id"`
	Type string `json:"type"`
	IsPublic bool `json:"isPublic"`
	Name string `json:"name"`
	Description string `json:"description"`
	Members []string `json:"members"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}
