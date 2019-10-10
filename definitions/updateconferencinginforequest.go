package definitions

// UpdateConferencingInfoRequest Update Conferencing Info Request
type UpdateConferencingInfoRequest struct {
	PhoneNumbers []ConferencePhoneNumberInfo `json:"phoneNumbers"`
	AllowJoinBeforeHost bool `json:"allowJoinBeforeHost"`
}
