package definitions

// UserCallLogRecord User Call Log Record
type UserCallLogRecord struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	SessionId string `json:"sessionId"`
	From CallLogCallerInfo `json:"from"`
	To CallLogCallerInfo `json:"to"`
	Type string `json:"type"`
	Direction string `json:"direction"`
	StartTime string `json:"startTime"`
	Deleted bool `json:"deleted"`
	Duration int `json:"duration"`
	Recording CallLogRecordingInfo `json:"recording"`
	Action string `json:"action"`
	Result string `json:"result"`
	Reason string `json:"reason"`
}
