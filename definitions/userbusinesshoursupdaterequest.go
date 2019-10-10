package definitions

// UserBusinessHoursUpdateRequest User Business Hours Update Request
type UserBusinessHoursUpdateRequest struct {
	Schedule UserBusinessHoursScheduleInfo `json:"schedule"`
}
