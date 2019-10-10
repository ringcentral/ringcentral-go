package definitions

// AccountPhoneNumbers Account Phone Numbers
type AccountPhoneNumbers struct {
	Records []CompanyPhoneNumberInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
