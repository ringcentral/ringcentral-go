package definitions

// CreateSipRegistrationResponse Create Sip Registration Response
type CreateSipRegistrationResponse struct {
	Device SipRegistrationDeviceInfo `json:"device"`
	SipInfo []SIPInfoResponse `json:"sipInfo"`
	SipInfoPstn []SIPInfoResponse `json:"sipInfoPstn"`
	SipFlags []SIPFlagsResponse `json:"sipFlags"`
	SipErrorCodes []string `json:"sipErrorCodes"`
}
