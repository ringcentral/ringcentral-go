package definitions

// RecordsCollectionResourceSubscriptionResponse Records Collection Resource Subscription Response
type RecordsCollectionResourceSubscriptionResponse struct {
	Uri string `json:"uri"`
	Records []SubscriptionInfo `json:"records"`
}
