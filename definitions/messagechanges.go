package definitions

// MessageChanges Message Changes
type MessageChanges struct {
	Type string `json:"type"`
	NewCount int `json:"newCount"`
	UpdatedCount int `json:"updatedCount"`
}
