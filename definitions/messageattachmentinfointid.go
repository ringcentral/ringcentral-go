package definitions

// MessageAttachmentInfoIntId Message Attachment Info Int Id
type MessageAttachmentInfoIntId struct {
	Id int `json:"id"`
	Uri string `json:"uri"`
	Type string `json:"type"`
	ContentType string `json:"contentType"`
	VmDuration int `json:"vmDuration"`
	Filename string `json:"filename"`
	Size int `json:"size"`
}
