package definitions

// CustomUserGreetingInfo Custom User Greeting Info
type CustomUserGreetingInfo struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	Type string `json:"type"`
	ContentType string `json:"contentType"`
	ContentUri string `json:"contentUri"`
	AnsweringRule CustomGreetingAnsweringRuleInfo `json:"answeringRule"`
}
