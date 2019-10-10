package definitions

// ForwardingNumberInfo Forwarding Number Info
type ForwardingNumberInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	PhoneNumber string `json:"phoneNumber"`
	Label string `json:"label"`
	Features []string `json:"features"`
	FlipNumber int `json:"flipNumber"`
	Type string `json:"type"`
}
