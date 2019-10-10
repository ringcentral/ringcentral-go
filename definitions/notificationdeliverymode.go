package definitions

// NotificationDeliveryMode Notification Delivery Mode
type NotificationDeliveryMode struct {
	TransportType string `json:"transportType"`
	Encryption bool `json:"encryption"`
	Address string `json:"address"`
	SubscriberKey string `json:"subscriberKey"`
	SecretKey string `json:"secretKey"`
	EncryptionAlgorithm string `json:"encryptionAlgorithm"`
	EncryptionKey string `json:"encryptionKey"`
}
