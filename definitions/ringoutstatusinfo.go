package definitions

// RingOutStatusInfo Ring Out Status Info
type RingOutStatusInfo struct {
	CallStatus string `json:"callStatus"`
	CallerStatus string `json:"callerStatus"`
	CalleeStatus string `json:"calleeStatus"`
}
