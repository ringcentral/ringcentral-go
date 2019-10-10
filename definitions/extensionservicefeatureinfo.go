package definitions

// ExtensionServiceFeatureInfo Extension Service Feature Info
type ExtensionServiceFeatureInfo struct {
	Enabled bool `json:"enabled"`
	FeatureName string `json:"featureName"`
	Reason string `json:"reason"`
}
