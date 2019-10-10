package definitions

// GetUserBusinessHoursResponse Get User Business Hours Response
type GetUserBusinessHoursResponse struct {
	Uri string `json:"uri"`
	Schedule ScheduleInfoUserBusinessHours `json:"schedule"`
}
