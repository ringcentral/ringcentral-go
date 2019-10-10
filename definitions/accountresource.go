package definitions

// AccountResource Account Resource
type AccountResource struct {
	CompanyName string `json:"companyName"`
	FederatedName string `json:"federatedName"`
	Id string `json:"id"`
	MainNumber PhoneNumberResource `json:"mainNumber"`
}
