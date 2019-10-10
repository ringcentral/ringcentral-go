package definitions

// GlipMessageAttachmentInfo Glip Message Attachment Info
type GlipMessageAttachmentInfo struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Fallback string `json:"fallback"`
	Intro string `json:"intro"`
	Author GlipMessageAttachmentAuthorInfo `json:"author"`
	Title string `json:"title"`
	Text string `json:"text"`
	ImageUri string `json:"imageUri"`
	ThumbnailUri string `json:"thumbnailUri"`
	Fields []GlipMessageAttachmentFieldsInfo `json:"fields"`
	Footnote GlipMessageAttachmentFootnoteInfo `json:"footnote"`
	CreatorId string `json:"creatorId"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	AllDay bool `json:"allDay"`
	Recurrence string `json:"recurrence"`
	EndingCondition string `json:"endingCondition"`
	EndingAfter int `json:"endingAfter"`
	EndingOn string `json:"endingOn"`
	Color string `json:"color"`
	Location string `json:"location"`
	Description string `json:"description"`
}
