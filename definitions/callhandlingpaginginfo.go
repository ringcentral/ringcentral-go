package definitions

// CallHandlingPagingInfo Call Handling Paging Info
type CallHandlingPagingInfo struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
	PageStart int `json:"pageStart"`
	PageEnd int `json:"pageEnd"`
	TotalPages int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}
