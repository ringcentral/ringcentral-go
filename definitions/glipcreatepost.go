package definitions

// GlipCreatePost Glip Create Post
type GlipCreatePost struct {
	Title string `json:"title"`
	Text string `json:"text"`
	GroupId string `json:"groupId"`
	Attachments []GlipMessageAttachmentInfoRequest `json:"attachments"`
}
