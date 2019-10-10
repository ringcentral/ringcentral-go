package definitions

// ActiveCallInfo Active Call Info
type ActiveCallInfo struct {
	Id string `json:"id"`
	Direction string `json:"direction"`
	From string `json:"from"`
	FromName string `json:"fromName"`
	To string `json:"to"`
	ToName string `json:"toName"`
	StartTime string `json:"startTime"`
	TelephonyStatus string `json:"telephonyStatus"`
	SipData DetailedCallInfo `json:"sipData"`
	SessionId string `json:"sessionId"`
	TerminationType string `json:"terminationType"`
}
