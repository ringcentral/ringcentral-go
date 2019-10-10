package definitions

// ExtensionResourceIntId Extension Resource Int Id
type ExtensionResourceIntId struct {
	Uri string `json:"uri"`
	Id int `json:"id"`
	PartnerId string `json:"partnerId"`
	ExtensionNumber string `json:"extensionNumber"`
	LoginName string `json:"loginName"`
	Contact ExtensionContactInfo `json:"contact"`
	References []Reference `json:"references"`
	Name string `json:"name"`
	Type string `json:"type"`
	Status string `json:"status"`
	StatusInfo StatusInfo `json:"statusInfo"`
	Departments []DepartmentResource `json:"departments"`
	ServiceFeatures []ServiceFeatureValue `json:"serviceFeatures"`
	RegionalSettings RegionalSettingsInfo `json:"regionalSettings"`
	SetupWizardState string `json:"setupWizardState"`
	Permissions ExtensionPermissionsResource `json:"permissions"`
	Password string `json:"password"`
	IvrPin string `json:"ivrPin"`
	ProfileImage ProfileImageResource `json:"profileImage"`
}
