package definitions

// ActiveCallInfoWithoutSIP Active Call Info Without SIP
type ActiveCallInfoWithoutSIP struct {
	Id string `json:"id"`
	Direction string `json:"direction"`
	From string `json:"from"`
	FromName string `json:"fromName"`
	To string `json:"to"`
	ToName string `json:"toName"`
	StartTime string `json:"startTime"`
	TelephonyStatus string `json:"telephonyStatus"`
	SessionId string `json:"sessionId"`
	TerminationType string `json:"terminationType"`
}
