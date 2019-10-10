package definitions

// MeetingsPagingInfo Meetings Paging Info
type MeetingsPagingInfo struct {
	Page int `json:"page"`
	TotalPages int `json:"totalPages"`
	PerPage int `json:"perPage"`
	TotalElements int `json:"totalElements"`
	PageStart int `json:"pageStart"`
	PageEnd int `json:"pageEnd"`
}
