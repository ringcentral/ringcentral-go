package definitions

// FaxResponse Fax Response
type FaxResponse struct {
	Id int `json:"id"`
	Uri string `json:"uri"`
	Type string `json:"type"`
	From CallerInfoFrom `json:"from"`
	To []CallerInfoTo `json:"to"`
	CreationTime string `json:"creationTime"`
	ReadStatus string `json:"readStatus"`
	Priority string `json:"priority"`
	Attachments []MessageAttachmentInfoIntId `json:"attachments"`
	Direction string `json:"direction"`
	Availability string `json:"availability"`
	MessageStatus string `json:"messageStatus"`
	FaxResolution string `json:"faxResolution"`
	FaxPageCount int `json:"faxPageCount"`
	LastModifiedTime string `json:"lastModifiedTime"`
	CoverIndex int `json:"coverIndex"`
	CoverPageText string `json:"coverPageText"`
}
