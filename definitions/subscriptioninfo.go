package definitions

// SubscriptionInfo Subscription Info
type SubscriptionInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	EventFilters []string `json:"eventFilters"`
	DisabledFilters []DisabledFilterInfo `json:"disabledFilters"`
	ExpirationTime string `json:"expirationTime"`
	ExpiresIn int `json:"expiresIn"`
	Status string `json:"status"`
	CreationTime string `json:"creationTime"`
	DeliveryMode NotificationDeliveryMode `json:"deliveryMode"`
	BlacklistedData NotificationBlacklistedData `json:"blacklistedData"`
}
