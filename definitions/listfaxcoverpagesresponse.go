package definitions

// ListFaxCoverPagesResponse List Fax Cover Pages Response
type ListFaxCoverPagesResponse struct {
	Uri string `json:"uri"`
	Records []FaxCoverPageInfo `json:"records"`
	Navigation MessagingNavigationInfo `json:"navigation"`
	Paging MessagingPagingInfo `json:"paging"`
}
