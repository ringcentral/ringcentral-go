package definitions

// GrantInfo Grant Info
type GrantInfo struct {
	Uri string `json:"uri"`
	Extension ExtensionInfoGrants `json:"extension"`
	CallPickup bool `json:"callPickup"`
	CallMonitoring bool `json:"callMonitoring"`
	CallOnBehalfOf bool `json:"callOnBehalfOf"`
	CallDelegation bool `json:"callDelegation"`
	GroupPaging bool `json:"groupPaging"`
}
