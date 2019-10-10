package definitions

// CreateMultipleSwitchesRequest Create Multiple Switches Request
type CreateMultipleSwitchesRequest struct {
	Records []CreateSwitchInfo `json:"records"`
}
