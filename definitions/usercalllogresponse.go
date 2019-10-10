package definitions

// UserCallLogResponse User Call Log Response
type UserCallLogResponse struct {
	Records []CallLogRecord `json:"records"`
	Navigation CallLogNavigationInfo `json:"navigation"`
	Paging CallLogPagingInfo `json:"paging"`
}
