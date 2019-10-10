package definitions

// CreateSubscriptionRequest Create Subscription Request
type CreateSubscriptionRequest struct {
	EventFilters []string `json:"eventFilters"`
	DeliveryMode NotificationDeliveryModeRequest `json:"deliveryMode"`
	ExpiresIn int `json:"expiresIn"`
}
