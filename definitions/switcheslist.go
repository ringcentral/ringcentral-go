package definitions

// SwitchesList Switches List
type SwitchesList struct {
	Records []SwitchInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
