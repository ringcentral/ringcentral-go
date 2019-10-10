package definitions

// ScheduleInfo Schedule Info
type ScheduleInfo struct {
	WeeklyRanges WeeklyScheduleInfo `json:"weeklyRanges"`
	Ranges []RangesInfo `json:"ranges"`
	Ref string `json:"ref"`
}
