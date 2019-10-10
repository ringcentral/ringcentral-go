package definitions

// BillingPlanInfo Billing Plan Info
type BillingPlanInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	DurationUnit string `json:"durationUnit"`
	Duration string `json:"duration"`
	Type string `json:"type"`
}
