package definitions

// GlipNoteInfo Glip Note Info
type GlipNoteInfo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	ChatIds []string `json:"chatIds"`
	Preview string `json:"preview"`
	Creator NoteCreatorInfo `json:"creator"`
	LastModifiedBy LastModifiedByInfo `json:"lastModifiedBy"`
	LockedBy LockedByInfo `json:"lockedBy"`
	Status string `json:"status"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Type string `json:"type"`
}
