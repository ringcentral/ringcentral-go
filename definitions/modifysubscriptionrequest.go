package definitions

// ModifySubscriptionRequest Modify Subscription Request
type ModifySubscriptionRequest struct {
	EventFilters []string `json:"eventFilters"`
	DeliveryMode NotificationDeliveryModeRequest `json:"deliveryMode"`
	ExpiresIn int `json:"expiresIn"`
}
