package definitions

// CallLogCallerInfo Call Log Caller Info
type CallLogCallerInfo struct {
	PhoneNumber string `json:"phoneNumber"`
	ExtensionNumber string `json:"extensionNumber"`
	ExtensionId string `json:"extensionId"`
	Location string `json:"location"`
	Name string `json:"name"`
	Device CallLogRecordDeviceInfo `json:"device"`
}
