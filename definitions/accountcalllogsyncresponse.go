package definitions

// AccountCallLogSyncResponse Account Call Log Sync Response
type AccountCallLogSyncResponse struct {
	Records []CompanyCallLogRecord `json:"records"`
	SyncInfo CompanyCallLogSyncInfo `json:"syncInfo"`
}
