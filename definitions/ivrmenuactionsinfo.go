package definitions

// IVRMenuActionsInfo IVRMenu Actions Info
type IVRMenuActionsInfo struct {
	Input string `json:"input"`
	Action string `json:"action"`
	Extension IVRMenuExtensionInfo `json:"extension"`
	PhoneNumber string `json:"phoneNumber"`
}
