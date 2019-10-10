package definitions

// CallRecordingExtensions Call Recording Extensions
type CallRecordingExtensions struct {
	Uri string `json:"uri"`
	Records []CallRecordingExtensionInfo `json:"records"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging CallHandlingPagingInfo `json:"paging"`
}
