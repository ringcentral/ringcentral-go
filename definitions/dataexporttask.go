package definitions

// DataExportTask Data Export Task
type DataExportTask struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Status string `json:"status"`
	Result []ExportTaskResultInfo `json:"result"`
}
