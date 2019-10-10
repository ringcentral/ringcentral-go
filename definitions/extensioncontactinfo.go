package definitions

// ExtensionContactInfo Extension Contact Info
type ExtensionContactInfo struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Company string `json:"company"`
	Email string `json:"email"`
	BusinessPhone string `json:"businessPhone"`
	BusinessAddress ContactAddressInfoDevices `json:"businessAddress"`
}
