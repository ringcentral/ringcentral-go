package definitions

// PrivateIpRangeInfoRequest Private Ip Range Info Request
type PrivateIpRangeInfoRequest struct {
	Id string `json:"id"`
	StartIp string `json:"startIp"`
	EndIp string `json:"endIp"`
	Name string `json:"name"`
	EmergencyAddress LocationUpdatesEmergencyAddressInfoRequest `json:"emergencyAddress"`
	EmergencyLocationId string `json:"emergencyLocationId"`
}
