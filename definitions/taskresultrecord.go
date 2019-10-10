package definitions

// TaskResultRecord Task Result Record
type TaskResultRecord struct {
	Id string `json:"id"`
	Bssid string `json:"bssid"`
	ChassisId string `json:"chassisId"`
	Status string `json:"status"`
	Errors TaskResultRecordErrorsInfo `json:"errors"`
}
