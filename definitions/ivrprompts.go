package definitions

// IVRPrompts IVRPrompts
type IVRPrompts struct {
	Uri string `json:"uri"`
	Records []PromptInfo `json:"records"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging CallHandlingPagingInfo `json:"paging"`
}
