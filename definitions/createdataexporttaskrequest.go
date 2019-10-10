package definitions

// CreateDataExportTaskRequest Create Data Export Task Request
type CreateDataExportTaskRequest struct {
	DateFrom string `json:"dateFrom"`
	DateTo string `json:"dateTo"`
	UserIds []string `json:"userIds"`
	ChatIds []string `json:"chatIds"`
}
