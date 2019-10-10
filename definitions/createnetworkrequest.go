package definitions

// CreateNetworkRequest Create Network Request
type CreateNetworkRequest struct {
	Name string `json:"name"`
	Site AutomaticLocationUpdatesSiteInfo `json:"site"`
	PublicIpRanges []PublicIpRangeInfo `json:"publicIpRanges"`
	PrivateIpRanges []PrivateIpRangeInfoRequest `json:"privateIpRanges"`
}
