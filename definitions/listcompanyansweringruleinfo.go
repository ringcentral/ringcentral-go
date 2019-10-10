package definitions

// ListCompanyAnsweringRuleInfo List Company Answering Rule Info
type ListCompanyAnsweringRuleInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Enabled bool `json:"enabled"`
	Type string `json:"type"`
	Name string `json:"name"`
}
