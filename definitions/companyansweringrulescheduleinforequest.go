package definitions

// CompanyAnsweringRuleScheduleInfoRequest Company Answering Rule Schedule Info Request
type CompanyAnsweringRuleScheduleInfoRequest struct {
	WeeklyRanges CompanyAnsweringRuleWeeklyScheduleInfoRequest `json:"weeklyRanges"`
	Ranges []RangesInfo `json:"ranges"`
	Ref string `json:"ref"`
}
