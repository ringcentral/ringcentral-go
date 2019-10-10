package definitions

// ContactInfoUpdateRequest Contact Info Update Request
type ContactInfoUpdateRequest struct {
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
