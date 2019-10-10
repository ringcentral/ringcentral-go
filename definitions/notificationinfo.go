package definitions

// NotificationInfo Notification Info
type NotificationInfo struct {
	Alert AlertInfo `json:"alert"`
	Badge string `json:"badge"`
	Sound string `json:"sound"`
	ContentAvailable string `json:"content-available"`
	Category string `json:"category"`
}
