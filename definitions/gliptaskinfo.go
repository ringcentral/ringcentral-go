package definitions

// GlipTaskInfo Glip Task Info
type GlipTaskInfo struct {
	Id string `json:"id"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
	Type string `json:"type"`
	Creator interface{} `json:"creator"`
	ChatIds []string `json:"chatIds"`
	Status string `json:"status"`
	Subject string `json:"subject"`
	Assignees []interface{} `json:"assignees"`
	CompletenessCondition string `json:"completenessCondition"`
	CompletenessPercentage int `json:"completenessPercentage"`
	StartDate string `json:"startDate"`
	DueDate string `json:"dueDate"`
	Color string `json:"color"`
	Section string `json:"section"`
	Description string `json:"description"`
	Recurrence GlipTaskRecurrenceInfo `json:"recurrence"`
	Attachments []TaskAttachment `json:"attachments"`
}
