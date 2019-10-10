package definitions

// WeeklyScheduleInfo Weekly Schedule Info
type WeeklyScheduleInfo struct {
	Monday []TimeInterval `json:"monday"`
	Tuesday []TimeInterval `json:"tuesday"`
	Wednesday []TimeInterval `json:"wednesday"`
	Thursday []TimeInterval `json:"thursday"`
	Friday []TimeInterval `json:"friday"`
	Saturday []TimeInterval `json:"saturday"`
	Sunday []TimeInterval `json:"sunday"`
}
