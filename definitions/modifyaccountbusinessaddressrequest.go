package definitions

// ModifyAccountBusinessAddressRequest Modify Account Business Address Request
type ModifyAccountBusinessAddressRequest struct {
	Company string `json:"company"`
	Email string `json:"email"`
	BusinessAddress BusinessAddressInfo `json:"businessAddress"`
}
