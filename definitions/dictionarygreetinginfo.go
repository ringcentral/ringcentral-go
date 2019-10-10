package definitions

// DictionaryGreetingInfo Dictionary Greeting Info
type DictionaryGreetingInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Name string `json:"name"`
	UsageType string `json:"usageType"`
	Text string `json:"text"`
	ContentUri string `json:"contentUri"`
	Type string `json:"type"`
	Category string `json:"category"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging CallHandlingPagingInfo `json:"paging"`
}
