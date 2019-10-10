package definitions

// ExtensionCreationRequest Extension Creation Request
type ExtensionCreationRequest struct {
	Contact ContactInfoCreationRequest `json:"contact"`
	ExtensionNumber string `json:"extensionNumber"`
	Password string `json:"password"`
	References []ReferenceInfo `json:"references"`
	Roles []Roles `json:"roles"`
	RegionalSettings RegionalSettings `json:"regionalSettings"`
	SetupWizardState string `json:"setupWizardState"`
	Status string `json:"status"`
	StatusInfo ExtensionStatusInfo `json:"statusInfo"`
	Type string `json:"type"`
	Hidden bool `json:"hidden"`
}
