package definitions

// NotificationRecipientInfo Notification Recipient Info
type NotificationRecipientInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	ExtensionNumber string `json:"extensionNumber"`
	Target bool `json:"target"`
	Location string `json:"location"`
	Name string `json:"name"`
}
