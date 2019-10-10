package definitions

// GlipCompleteTask Glip Complete Task
type GlipCompleteTask struct {
	Status string `json:"status"`
	Assignees []interface{} `json:"assignees"`
	CompletenessPercentage int `json:"completenessPercentage"`
}
