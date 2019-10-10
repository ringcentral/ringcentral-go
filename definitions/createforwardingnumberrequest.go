package definitions

// CreateForwardingNumberRequest Create Forwarding Number Request
type CreateForwardingNumberRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Label string `json:"label"`
	Type string `json:"type"`
	Device CreateForwardingNumberDeviceInfo `json:"device"`
}
