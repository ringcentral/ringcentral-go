package definitions

// CompanyContactsPagingInfo Company Contacts Paging Info
type CompanyContactsPagingInfo struct {
	Page int `json:"page"`
	TotalPages int `json:"totalPages"`
	PerPage int `json:"perPage"`
	TotalElements int `json:"totalElements"`
	PageStart int `json:"pageStart"`
	PageEnd int `json:"pageEnd"`
}
