package definitions

// ActiveCallsResponse Active Calls Response
type ActiveCallsResponse struct {
	Records []CallLogRecord `json:"records"`
	Navigation CallLogNavigationInfo `json:"navigation"`
	Paging CallLogPagingInfo `json:"paging"`
}
