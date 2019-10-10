package definitions

// IVRMenuPromptInfo IVRMenu Prompt Info
type IVRMenuPromptInfo struct {
	Mode string `json:"mode"`
	Audio PromptLanguageInfo `json:"audio"`
	Text string `json:"text"`
	Language AudioPromptInfo `json:"language"`
}
