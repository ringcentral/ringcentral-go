package definitions

// GlipPostEvent Glip Post Event
type GlipPostEvent struct {
	Id string `json:"id"`
	EventType string `json:"eventType"`
	GroupId string `json:"groupId"`
	Type string `json:"type"`
	Text string `json:"text"`
	CreatorId string `json:"creatorId"`
	AddedPersonIds []string `json:"addedPersonIds"`
	RemovedPersonIds []string `json:"removedPersonIds"`
	Mentions []GlipMentionsInfo `json:"mentions"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}
