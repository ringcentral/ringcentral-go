package definitions

// MeetingServiceInfoResource Meeting Service Info Resource
type MeetingServiceInfoResource struct {
	Uri string `json:"uri"`
	SupportUri string `json:"supportUri"`
	IntlDialInNumbersUri string `json:"intlDialInNumbersUri"`
	ExternalUserInfo MeetingExternalUserInfoResource `json:"externalUserInfo"`
	DialInNumbers []DialInNumberResource `json:"dialInNumbers"`
}
