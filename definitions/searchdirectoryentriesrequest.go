package definitions

// SearchDirectoryEntriesRequest Search Directory Entries Request
type SearchDirectoryEntriesRequest struct {
	SearchString string `json:"searchString"`
	ShowFederated bool `json:"showFederated"`
	ExtensionType string `json:"extensionType"`
	OrderBy []OrderBy `json:"orderBy"`
	Page int `json:"page"`
	PerPage int `json:"perPage"`
}
