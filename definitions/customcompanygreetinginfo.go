package definitions

// CustomCompanyGreetingInfo Custom Company Greeting Info
type CustomCompanyGreetingInfo struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	Type string `json:"type"`
	ContentType string `json:"contentType"`
	ContentUri string `json:"contentUri"`
	AnsweringRule CustomGreetingAnsweringRuleInfo `json:"answeringRule"`
	Language CustomCompanyGreetingLanguageInfo `json:"language"`
}
