package definitions

// ContactResource Contact Resource
type ContactResource struct {
	Account AccountResource `json:"account"`
	Department string `json:"department"`
	Email string `json:"email"`
	ExtensionNumber string `json:"extensionNumber"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Name string `json:"name"`
	Id string `json:"id"`
	JobTitle string `json:"jobTitle"`
	PhoneNumbers []PhoneNumberResource `json:"phoneNumbers"`
	ProfileImage AccountDirectoryProfileImageResource `json:"profileImage"`
	Site BusinessSiteResource `json:"site"`
	Status string `json:"status"`
	Type string `json:"type"`
}
