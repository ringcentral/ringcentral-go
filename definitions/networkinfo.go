package definitions

// NetworkInfo Network Info
type NetworkInfo struct {
	Name string `json:"name"`
	Site AutomaticLocationUpdatesSiteInfo `json:"site"`
	PublicIpRanges []PublicIpRangeInfo `json:"publicIpRanges"`
	PrivateIpRanges []PrivateIpRangeInfo `json:"privateIpRanges"`
}
