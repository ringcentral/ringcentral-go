package definitions

// MissedCallsInfo Missed Calls Info
type MissedCallsInfo struct {
	NotifyByEmail bool `json:"notifyByEmail"`
	NotifyBySms bool `json:"notifyBySms"`
	AdvancedEmailAddresses []string `json:"advancedEmailAddresses"`
	AdvancedSmsEmailAddresses []string `json:"advancedSmsEmailAddresses"`
}
