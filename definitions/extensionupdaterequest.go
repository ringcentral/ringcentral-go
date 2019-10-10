package definitions

// ExtensionUpdateRequest Extension Update Request
type ExtensionUpdateRequest struct {
	Status string `json:"status"`
	StatusInfo ExtensionStatusInfo `json:"statusInfo"`
	Reason string `json:"reason"`
	Comment string `json:"comment"`
	ExtensionNumber string `json:"extensionNumber"`
	Contact ContactInfoUpdateRequest `json:"contact"`
	RegionalSettings ExtensionRegionalSettingRequest `json:"regionalSettings"`
	SetupWizardState string `json:"setupWizardState"`
	PartnerId string `json:"partnerId"`
	IvrPin string `json:"ivrPin"`
	Password string `json:"password"`
	CallQueueInfo CallQueueInfoRequest `json:"callQueueInfo"`
	Transition []UserTransitionInfo `json:"transition"`
}
