package definitions

// CreateMultipleWirelessPointsRequest Create Multiple Wireless Points Request
type CreateMultipleWirelessPointsRequest struct {
	Records []CreateWirelessPoint `json:"records"`
}
