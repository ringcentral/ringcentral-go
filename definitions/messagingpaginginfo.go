package definitions

// MessagingPagingInfo Messaging Paging Info
type MessagingPagingInfo struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
	PageStart int `json:"pageStart"`
	PageEnd int `json:"pageEnd"`
	TotalPages int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}
