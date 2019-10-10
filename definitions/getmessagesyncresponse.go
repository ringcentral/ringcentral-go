package definitions

// GetMessageSyncResponse Get Message Sync Response
type GetMessageSyncResponse struct {
	Records []GetMessageInfoResponse `json:"records"`
	SyncInfo SyncInfoMessages `json:"syncInfo"`
}
