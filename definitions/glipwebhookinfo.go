package definitions

// GlipWebhookInfo Glip Webhook Info
type GlipWebhookInfo struct {
	Id string `json:"id"`
	CreatorId string `json:"creatorId"`
	GroupIds []string `json:"groupIds"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Uri string `json:"uri"`
	Status string `json:"status"`
}
