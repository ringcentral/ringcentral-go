package definitions

// ContactInfoCreationRequest Contact Info Creation Request
type ContactInfoCreationRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Company string `json:"company"`
	JobTitle string `json:"jobTitle"`
	Email string `json:"email"`
	BusinessPhone string `json:"businessPhone"`
	MobilePhone string `json:"mobilePhone"`
	BusinessAddress ContactBusinessAddressInfo `json:"businessAddress"`
	EmailAsLoginName bool `json:"emailAsLoginName"`
	PronouncedName PronouncedNameInfo `json:"pronouncedName"`
	Department string `json:"department"`
}
