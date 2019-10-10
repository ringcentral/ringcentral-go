package definitions

// GlipNotesInfo Glip Notes Info
type GlipNotesInfo struct {
	Records []GlipNoteInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
