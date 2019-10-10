package definitions

// UserAnsweringRuleListNavigation User Answering Rule List Navigation
type UserAnsweringRuleListNavigation struct {
	FirstPage UserAnsweringRuleListNavigationPage `json:"firstPage"`
	NextPage UserAnsweringRuleListNavigationPage `json:"nextPage"`
	PreviousPage UserAnsweringRuleListNavigationPage `json:"previousPage"`
	LastPage UserAnsweringRuleListNavigationPage `json:"lastPage"`
}
