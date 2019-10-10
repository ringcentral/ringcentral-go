package definitions

// UserTemplates User Templates
type UserTemplates struct {
	Uri string `json:"uri"`
	Records []TemplateInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
