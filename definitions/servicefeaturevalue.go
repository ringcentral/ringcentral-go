package definitions

// ServiceFeatureValue Service Feature Value
type ServiceFeatureValue struct {
	FeatureName string `json:"featureName"`
	Enabled bool `json:"enabled"`
	Reason string `json:"reason"`
}
