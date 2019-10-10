package definitions

// UserResponse User Response
type UserResponse struct {
	Active bool `json:"active"`
	Addresses []UserAddress `json:"addresses"`
	Emails []Email `json:"emails"`
	ExternalId string `json:"externalId"`
	Id string `json:"id"`
	Name Name `json:"name"`
	PhoneNumbers []PhoneNumber `json:"phoneNumbers"`
	Photos []Photo `json:"photos"`
	Schemas []string `json:"schemas"`
	UrnIetfParamsScimSchemasExtensionEnterprise_2_0User EnterpriseUser `json:"urn:ietf:params:scim:schemas:extension:enterprise:2.0:User"`
	UserName string `json:"userName"`
	Meta Meta `json:"meta"`
}
