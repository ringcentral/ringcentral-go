package definitions

// IVRMenuInfo IVRMenu Info
type IVRMenuInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Name string `json:"name"`
	ExtensionNumber string `json:"extensionNumber"`
	Prompt IVRMenuPromptInfo `json:"prompt"`
	Actions []IVRMenuActionsInfo `json:"actions"`
}
