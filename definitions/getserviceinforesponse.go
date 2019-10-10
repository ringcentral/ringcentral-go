package definitions

// GetServiceInfoResponse Get Service Info Response
type GetServiceInfoResponse struct {
	Uri string `json:"uri"`
	ServicePlanName string `json:"servicePlanName"`
	ServiceFeatures []ServiceFeatureInfo `json:"serviceFeatures"`
	Limits AccountLimits `json:"limits"`
}
