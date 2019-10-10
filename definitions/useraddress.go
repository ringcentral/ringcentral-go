package definitions

// UserAddress User Address
type UserAddress struct {
	Country string `json:"country"`
	Locality string `json:"locality"`
	PostalCode string `json:"postalCode"`
	Region string `json:"region"`
	StreetAddress string `json:"streetAddress"`
	Type string `json:"type"`
}
