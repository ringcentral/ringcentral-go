package definitions

// GetConferencingInfoResponse Get Conferencing Info Response
type GetConferencingInfoResponse struct {
	Uri string `json:"uri"`
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost"`
	HostCode string `json:"hostCode"`
	Mode string `json:"mode"`
	ParticipantCode string `json:"participantCode"`
	PhoneNumber string `json:"phoneNumber"`
	TapToJoinUri string `json:"tapToJoinUri"`
	PhoneNumbers []PhoneNumberInfoConferencing `json:"phoneNumbers"`
}
