package definitions

// CompanyAnsweringRuleWeeklyScheduleInfoRequest Company Answering Rule Weekly Schedule Info Request
type CompanyAnsweringRuleWeeklyScheduleInfoRequest struct {
	Monday []CompanyAnsweringRuleTimeIntervalRequest `json:"monday"`
	Tuesday []CompanyAnsweringRuleTimeIntervalRequest `json:"tuesday"`
	Wednesday []CompanyAnsweringRuleTimeIntervalRequest `json:"wednesday"`
	Thursday []CompanyAnsweringRuleTimeIntervalRequest `json:"thursday"`
	Friday []CompanyAnsweringRuleTimeIntervalRequest `json:"friday"`
	Saturday []CompanyAnsweringRuleTimeIntervalRequest `json:"saturday"`
	Sunday []CompanyAnsweringRuleTimeIntervalRequest `json:"sunday"`
}
