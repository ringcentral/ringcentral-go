package definitions

// CreateSipRegistrationRequest Create Sip Registration Request
type CreateSipRegistrationRequest struct {
	Device DeviceInfoRequest `json:"device"`
	SipInfo []SIPInfoRequest `json:"sipInfo"`
}
