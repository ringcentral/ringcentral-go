package definitions

// FederatedAccountResource Federated Account Resource
type FederatedAccountResource struct {
	CompanyName string `json:"companyName"`
	ConflictCount int `json:"conflictCount"`
	FederatedName string `json:"federatedName"`
	Id string `json:"id"`
	LinkCreationTime string `json:"linkCreationTime"`
	MainNumber PhoneNumberResource `json:"mainNumber"`
}
