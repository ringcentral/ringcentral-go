package definitions

// UpdateMultipleWirelessPointsRequest Update Multiple Wireless Points Request
type UpdateMultipleWirelessPointsRequest struct {
	Records []UpdateWirelessPoint `json:"records"`
}
