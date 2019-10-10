package definitions

// GlipCreateTask Glip Create Task
type GlipCreateTask struct {
	Subject string `json:"subject"`
	Assignees []interface{} `json:"assignees"`
	CompletenessCondition string `json:"completenessCondition"`
	StartDate string `json:"startDate"`
	DueDate string `json:"dueDate"`
	Color string `json:"color"`
	Section string `json:"section"`
	Description string `json:"description"`
	Recurrence GlipTaskRecurrenceInfo `json:"recurrence"`
	Attachments []interface{} `json:"attachments"`
}
