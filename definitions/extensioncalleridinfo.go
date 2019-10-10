package definitions

// ExtensionCallerIdInfo Extension Caller Id Info
type ExtensionCallerIdInfo struct {
	Uri string `json:"uri"`
	ByDevice []CallerIdByDevice `json:"byDevice"`
	ByFeature []CallerIdByFeature `json:"byFeature"`
	ExtensionNameForOutboundCalls bool `json:"extensionNameForOutboundCalls"`
	ExtensionNumberForInternalCalls bool `json:"extensionNumberForInternalCalls"`
}
