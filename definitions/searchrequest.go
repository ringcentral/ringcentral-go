package definitions

// SearchRequest Search Request
type SearchRequest struct {
	Count int `json:"count"`
	Filter string `json:"filter"`
	Schemas []string `json:"schemas"`
	StartIndex int `json:"startIndex"`
}
