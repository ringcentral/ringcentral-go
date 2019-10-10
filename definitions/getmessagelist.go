package definitions

// GetMessageList Get Message List
type GetMessageList struct {
	Records []GetMessageInfoResponse `json:"records"`
	Navigation MessagingNavigationInfo `json:"navigation"`
	Paging MessagingPagingInfo `json:"paging"`
}
