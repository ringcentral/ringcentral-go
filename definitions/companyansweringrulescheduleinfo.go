package definitions

// CompanyAnsweringRuleScheduleInfo Company Answering Rule Schedule Info
type CompanyAnsweringRuleScheduleInfo struct {
	WeeklyRanges CompanyAnsweringRuleWeeklyScheduleInfoRequest `json:"weeklyRanges"`
	Ranges []RangesInfo `json:"ranges"`
	Ref string `json:"ref"`
}
