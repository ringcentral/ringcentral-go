package definitions

// SyncInfo Sync Info
type SyncInfo struct {
	SyncType string `json:"syncType"`
	SyncToken string `json:"syncToken"`
	SyncTime string `json:"syncTime"`
	OlderRecordsExist bool `json:"olderRecordsExist"`
}
