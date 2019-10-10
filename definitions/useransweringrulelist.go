package definitions

// UserAnsweringRuleList User Answering Rule List
type UserAnsweringRuleList struct {
	Uri string `json:"uri"`
	Records []UserAnsweringRuleListRecord `json:"records"`
	Paging UserAnsweringRuleListPaging `json:"paging"`
	Navigation UserAnsweringRuleListNavigation `json:"navigation"`
}
