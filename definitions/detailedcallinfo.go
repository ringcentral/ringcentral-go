package definitions

// DetailedCallInfo Detailed Call Info
type DetailedCallInfo struct {
	CallId string `json:"callId"`
	ToTag string `json:"toTag"`
	FromTag string `json:"fromTag"`
	RemoteUri string `json:"remoteUri"`
	LocalUri string `json:"localUri"`
	RcSessionId string `json:"rcSessionId"`
}
