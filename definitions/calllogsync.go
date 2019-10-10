package definitions

// CallLogSync Call Log Sync
type CallLogSync struct {
	Records []CallLogRecord `json:"records"`
	SyncInfo SyncInfoCallLog `json:"syncInfo"`
}
