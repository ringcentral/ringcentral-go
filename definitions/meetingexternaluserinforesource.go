package definitions

// MeetingExternalUserInfoResource Meeting External User Info Resource
type MeetingExternalUserInfoResource struct {
	Uri string `json:"uri"`
	UserId string `json:"userId"`
	AccountId string `json:"accountId"`
	UserType int `json:"userType"`
	UserToken string `json:"userToken"`
	HostKey string `json:"hostKey"`
	PersonalMeetingId string `json:"personalMeetingId"`
}
