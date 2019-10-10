package definitions

// FederationResource Federation Resource
type FederationResource struct {
	Accounts []FederatedAccountResource `json:"accounts"`
	CreationTime string `json:"creationTime"`
	DisplayName string `json:"displayName"`
	Id string `json:"id"`
	LastModifiedTime string `json:"lastModifiedTime"`
}
