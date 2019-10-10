package definitions

// UpdateForwardingNumberRequest Update Forwarding Number Request
type UpdateForwardingNumberRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Label string `json:"label"`
	FlipNumber string `json:"flipNumber"`
	Type string `json:"type"`
}
