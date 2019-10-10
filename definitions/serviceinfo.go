package definitions

// ServiceInfo Service Info
type ServiceInfo struct {
	Uri string `json:"uri"`
	BillingPlan BillingPlanInfo `json:"billingPlan"`
	Brand BrandInfo `json:"brand"`
	ServicePlan ServicePlanInfo `json:"servicePlan"`
	TargetServicePlan TargetServicePlanInfo `json:"targetServicePlan"`
}
