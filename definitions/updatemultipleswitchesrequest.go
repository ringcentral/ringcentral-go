package definitions

// UpdateMultipleSwitchesRequest Update Multiple Switches Request
type UpdateMultipleSwitchesRequest struct {
	Records []UpdateSwitchInfo `json:"records"`
}
