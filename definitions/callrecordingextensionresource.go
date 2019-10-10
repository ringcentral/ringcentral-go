package definitions

// CallRecordingExtensionResource Call Recording Extension Resource
type CallRecordingExtensionResource struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	ExtensionNumber string `json:"extensionNumber"`
	Type string `json:"type"`
	CallDirection string `json:"callDirection"`
}
