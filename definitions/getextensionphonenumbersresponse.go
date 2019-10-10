package definitions

// GetExtensionPhoneNumbersResponse Get Extension Phone Numbers Response
type GetExtensionPhoneNumbersResponse struct {
	Records []UserPhoneNumberInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
