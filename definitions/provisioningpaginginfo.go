package definitions

// ProvisioningPagingInfo Provisioning Paging Info
type ProvisioningPagingInfo struct {
	Page int `json:"page"`
	PerPage int `json:"perPage"`
	PageStart int `json:"pageStart"`
	PageEnd int `json:"pageEnd"`
	TotalPages int `json:"totalPages"`
	TotalElements int `json:"totalElements"`
}
