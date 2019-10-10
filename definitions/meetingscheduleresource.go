package definitions

// MeetingScheduleResource Meeting Schedule Resource
type MeetingScheduleResource struct {
	StartTime string `json:"startTime"`
	DurationInMinutes int `json:"durationInMinutes"`
	TimeZone TimezoneResource `json:"timeZone"`
}
