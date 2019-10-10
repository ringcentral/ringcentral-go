package definitions

// UpdateNetworkRequest Update Network Request
type UpdateNetworkRequest struct {
	Name string `json:"name"`
	Site AutomaticLocationUpdatesSiteInfo `json:"site"`
	PublicIpRanges []PublicIpRangeInfo `json:"publicIpRanges"`
	PrivateIpRanges []PrivateIpRangeInfoRequest `json:"privateIpRanges"`
}
