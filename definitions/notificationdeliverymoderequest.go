package definitions

// NotificationDeliveryModeRequest Notification Delivery Mode Request
type NotificationDeliveryModeRequest struct {
	TransportType string `json:"transportType"`
	Address string `json:"address"`
	Encryption bool `json:"encryption"`
	CertificateName string `json:"certificateName"`
	RegistrationId string `json:"registrationId"`
	VerificationToken string `json:"verificationToken"`
}
