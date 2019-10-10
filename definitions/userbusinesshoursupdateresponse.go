package definitions

// UserBusinessHoursUpdateResponse User Business Hours Update Response
type UserBusinessHoursUpdateResponse struct {
	Uri string `json:"uri"`
	Schedule UserBusinessHoursScheduleInfo `json:"schedule"`
}
