package definitions

// CompanyAnsweringRuleList Company Answering Rule List
type CompanyAnsweringRuleList struct {
	Uri string `json:"uri"`
	Records []ListCompanyAnsweringRuleInfo `json:"records"`
	Paging CallHandlingPagingInfo `json:"paging"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
}
