package definitions

// BulkAssignItem Bulk Assign Item
type BulkAssignItem struct {
	DepartmentId string `json:"departmentId"`
	AddedExtensionIds []string `json:"addedExtensionIds"`
	RemovedExtensionIds []string `json:"removedExtensionIds"`
}
