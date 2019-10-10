package definitions

// MessageAttachmentInfo Message Attachment Info
type MessageAttachmentInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Type string `json:"type"`
	ContentType string `json:"contentType"`
	VmDuration int `json:"vmDuration"`
	Filename string `json:"filename"`
	Size int `json:"size"`
}
