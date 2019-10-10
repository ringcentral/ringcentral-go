package definitions

// MakeCallOutRequest Make Call Out Request
type MakeCallOutRequest struct {
	From MakeCallOutCallerInfoRequestFrom `json:"from"`
	To MakeCallOutCallerInfoRequestTo `json:"to"`
}
