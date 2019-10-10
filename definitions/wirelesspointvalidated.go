package definitions

// WirelessPointValidated Wireless Point Validated
type WirelessPointValidated struct {
	Id string `json:"id"`
	Bssid string `json:"bssid"`
	Status string `json:"status"`
	Errors []ValidationError `json:"errors"`
}
