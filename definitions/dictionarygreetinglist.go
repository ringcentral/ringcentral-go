package definitions

// DictionaryGreetingList Dictionary Greeting List
type DictionaryGreetingList struct {
	Uri string `json:"uri"`
	Records []DictionaryGreetingInfo `json:"records"`
}
