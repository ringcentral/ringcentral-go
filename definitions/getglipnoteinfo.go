package definitions

// GetGlipNoteInfo Get Glip Note Info
type GetGlipNoteInfo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	ChatIds []string `json:"chatIds"`
	Preview string `json:"preview"`
	Body string `json:"body"`
	Creator NoteCreatorInfo `json:"creator"`
	LastModifiedBy LastModifiedByInfo `json:"lastModifiedBy"`
	LockedBy LockedByInfo `json:"lockedBy"`
	Status string `json:"status"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Type string `json:"type"`
}
