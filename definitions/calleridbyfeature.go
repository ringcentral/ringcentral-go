package definitions

// CallerIdByFeature Caller Id By Feature
type CallerIdByFeature struct {
	Feature string `json:"feature"`
	CallerId CallerIdByFeatureInfo `json:"callerId"`
}
