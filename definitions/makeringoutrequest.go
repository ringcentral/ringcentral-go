package definitions

// MakeRingOutRequest Make Ring Out Request
type MakeRingOutRequest struct {
	From MakeRingOutCallerInfoRequestFrom `json:"from"`
	To MakeRingOutCallerInfoRequestTo `json:"to"`
	CallerId MakeRingOutCallerInfoRequestTo `json:"callerId"`
	PlayPrompt bool `json:"playPrompt"`
	Country MakeRingOutCoutryInfo `json:"country"`
}
