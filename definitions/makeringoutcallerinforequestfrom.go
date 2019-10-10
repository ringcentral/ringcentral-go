package definitions

// MakeRingOutCallerInfoRequestFrom Make Ring Out Caller Info Request From
type MakeRingOutCallerInfoRequestFrom struct {
	PhoneNumber string `json:"phoneNumber"`
	ForwardingNumberId string `json:"forwardingNumberId"`
}
