package definitions

// GetExtensionInfoResponse Get Extension Info Response
type GetExtensionInfoResponse struct {
	Id int `json:"id"`
	Uri string `json:"uri"`
	Contact ContactInfo `json:"contact"`
	Departments []DepartmentInfo `json:"departments"`
	ExtensionNumber string `json:"extensionNumber"`
	Name string `json:"name"`
	PartnerId string `json:"partnerId"`
	Permissions ExtensionPermissions `json:"permissions"`
	ProfileImage ProfileImageInfo `json:"profileImage"`
	References []ReferenceInfo `json:"references"`
	Roles []Roles `json:"roles"`
	RegionalSettings RegionalSettings `json:"regionalSettings"`
	ServiceFeatures []ExtensionServiceFeatureInfo `json:"serviceFeatures"`
	SetupWizardState string `json:"setupWizardState"`
	Status string `json:"status"`
	StatusInfo ExtensionStatusInfo `json:"statusInfo"`
	Type string `json:"type"`
	CallQueueExtensionInfo CallQueueExtensionInfo `json:"callQueueExtensionInfo"`
	Hidden bool `json:"hidden"`
}
