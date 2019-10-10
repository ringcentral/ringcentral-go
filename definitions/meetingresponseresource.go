package definitions

// MeetingResponseResource Meeting Response Resource
type MeetingResponseResource struct {
	Uri string `json:"uri"`
	Uuid string `json:"uuid"`
	Id string `json:"id"`
	Topic string `json:"topic"`
	MeetingType string `json:"meetingType"`
	Password string `json:"password"`
	H323Password string `json:"h323Password"`
	Status string `json:"status"`
	Links MeetingLinks `json:"links"`
	Schedule MeetingScheduleResource `json:"schedule"`
	Host HostInfoRequest `json:"host"`
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost"`
	StartHostVideo bool `json:"startHostVideo"`
	StartParticipantsVideo bool `json:"startParticipantsVideo"`
	AudioOptions []string `json:"audioOptions"`
}
