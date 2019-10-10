package definitions

// CompanyBusinessHours Company Business Hours
type CompanyBusinessHours struct {
	Uri string `json:"uri"`
	Schedule CompanyBusinessHoursScheduleInfo `json:"schedule"`
}
