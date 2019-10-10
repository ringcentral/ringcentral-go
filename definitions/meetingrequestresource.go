package definitions

// MeetingRequestResource Meeting Request Resource
type MeetingRequestResource struct {
	Topic string `json:"topic"`
	MeetingType string `json:"meetingType"`
	Schedule MeetingScheduleResource `json:"schedule"`
	Password string `json:"password"`
	Host HostInfoRequest `json:"host"`
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost"`
	StartHostVideo bool `json:"startHostVideo"`
	StartParticipantsVideo bool `json:"startParticipantsVideo"`
	UsePersonalMeetingId bool `json:"usePersonalMeetingId"`
	AudioOptions []string `json:"audioOptions"`
}
