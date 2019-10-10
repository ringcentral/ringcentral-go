package definitions

// CallLogRecord Call Log Record
type CallLogRecord struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	SessionId string `json:"sessionId"`
	From CallLogCallerInfo `json:"from"`
	To CallLogCallerInfo `json:"to"`
	Type string `json:"type"`
	Direction string `json:"direction"`
	Action string `json:"action"`
	Result string `json:"result"`
	Reason string `json:"reason"`
	StartTime string `json:"startTime"`
	Duration int `json:"duration"`
	Recording CallLogRecordingInfo `json:"recording"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Transport string `json:"transport"`
	Extension ActiveCallsRecordExtensionInfo `json:"extension"`
	Delegate DelegateInfo `json:"delegate"`
	Legs []CallLogRecordLegInfo `json:"legs"`
	Message CallLogRecordMessage `json:"message"`
	Deleted bool `json:"deleted"`
}
