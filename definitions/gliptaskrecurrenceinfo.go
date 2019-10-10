package definitions

// GlipTaskRecurrenceInfo Glip Task Recurrence Info
type GlipTaskRecurrenceInfo struct {
	Schedule string `json:"schedule"`
	EndingCondition string `json:"endingCondition"`
	EndingAfter int `json:"endingAfter"`
	EndingOn string `json:"endingOn"`
}
