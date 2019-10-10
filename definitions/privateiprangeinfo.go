package definitions

// PrivateIpRangeInfo Private Ip Range Info
type PrivateIpRangeInfo struct {
	Id string `json:"id"`
	StartIp string `json:"startIp"`
	EndIp string `json:"endIp"`
	Name string `json:"name"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfo `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
