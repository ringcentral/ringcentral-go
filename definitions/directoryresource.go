package definitions

// DirectoryResource Directory Resource
type DirectoryResource struct {
	Paging CompanyContactsPagingInfo `json:"paging"`
	Records []ContactResource `json:"records"`
}
