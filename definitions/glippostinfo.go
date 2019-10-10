package definitions

// GlipPostInfo Glip Post Info
type GlipPostInfo struct {
	Id string `json:"id"`
	GroupId string `json:"groupId"`
	Type string `json:"type"`
	Text string `json:"text"`
	CreatorId string `json:"creatorId"`
	AddedPersonIds []string `json:"addedPersonIds"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Attachments []GlipMessageAttachmentInfo `json:"attachments"`
	Mentions []GlipMentionsInfo `json:"mentions"`
	Activity string `json:"activity"`
	Title string `json:"title"`
	IconUri string `json:"iconUri"`
	IconEmoji string `json:"iconEmoji"`
}
