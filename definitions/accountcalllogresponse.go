package definitions

// AccountCallLogResponse Account Call Log Response
type AccountCallLogResponse struct {
	Records []CallLogRecord `json:"records"`
	Navigation CallLogNavigationInfo `json:"navigation"`
	Paging CallLogPagingInfo `json:"paging"`
}
