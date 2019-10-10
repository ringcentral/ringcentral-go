package definitions

// GlipMessageAttachmentInfoRequest Glip Message Attachment Info Request
type GlipMessageAttachmentInfoRequest struct {
	Type string `json:"type"`
	Title string `json:"title"`
	Fallback string `json:"fallback"`
	Color string `json:"color"`
	Intro string `json:"intro"`
	Author GlipMessageAttachmentAuthorInfo `json:"author"`
	Text string `json:"text"`
	ImageUri string `json:"imageUri"`
	ThumbnailUri string `json:"thumbnailUri"`
	Fields []GlipMessageAttachmentFieldsInfo `json:"fields"`
	Footnote GlipMessageAttachmentFootnoteInfo `json:"footnote"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	AllDay bool `json:"allDay"`
	Recurrence string `json:"recurrence"`
	EndingCondition string `json:"endingCondition"`
}
