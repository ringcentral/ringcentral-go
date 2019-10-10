package definitions

// CallQueueBulkAssignResource Call Queue Bulk Assign Resource
type CallQueueBulkAssignResource struct {
	AddedExtensionIds []string `json:"addedExtensionIds"`
	RemovedExtensionIds []string `json:"removedExtensionIds"`
}
