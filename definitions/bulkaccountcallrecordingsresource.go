package definitions

// BulkAccountCallRecordingsResource Bulk Account Call Recordings Resource
type BulkAccountCallRecordingsResource struct {
	AddedExtensions CallRecordingExtensionResource `json:"addedExtensions"`
	UpdatedExtensions CallRecordingExtensionResource `json:"updatedExtensions"`
	RemovedExtensions CallRecordingExtensionResource `json:"removedExtensions"`
}
