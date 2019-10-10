package definitions

// ForwardingNumberResource Forwarding Number Resource
type ForwardingNumberResource struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	PhoneNumber string `json:"phoneNumber"`
	Label string `json:"label"`
	Features []string `json:"features"`
	FlipNumber string `json:"flipNumber"`
	Type string `json:"type"`
}
