package definitions

// CompanyCallLogRecord Company Call Log Record
type CompanyCallLogRecord struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	SessionId string `json:"sessionId"`
	From CallLogCallerInfo `json:"from"`
	To CallLogCallerInfo `json:"to"`
	Type string `json:"type"`
	Direction string `json:"direction"`
	Deleted bool `json:"deleted"`
	Action string `json:"action"`
	Result string `json:"result"`
	Reason string `json:"reason"`
	StartTime string `json:"startTime"`
	Duration int `json:"duration"`
	Recording CallLogRecordingInfo `json:"recording"`
}
