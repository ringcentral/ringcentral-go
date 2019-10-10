package definitions

// LanguageList Language List
type LanguageList struct {
	Uri string `json:"uri"`
	Records []LanguageInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
