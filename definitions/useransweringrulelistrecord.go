package definitions

// UserAnsweringRuleListRecord User Answering Rule List Record
type UserAnsweringRuleListRecord struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Enabled bool `json:"enabled"`
}
