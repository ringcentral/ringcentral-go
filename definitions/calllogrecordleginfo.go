package definitions

// CallLogRecordLegInfo Call Log Record Leg Info
type CallLogRecordLegInfo struct {
	Action string `json:"action"`
	Direction string `json:"direction"`
	Duration int `json:"duration"`
	Extension ExtensionInfoCallLog `json:"extension"`
	LegType string `json:"legType"`
	StartTime string `json:"startTime"`
	Type string `json:"type"`
	Result string `json:"result"`
	Reason string `json:"reason"`
	From CallLogCallerInfo `json:"from"`
	To CallLogCallerInfo `json:"to"`
	Transport string `json:"transport"`
	Recording CallLogRecordingInfo `json:"recording"`
	Master bool `json:"master"`
}
