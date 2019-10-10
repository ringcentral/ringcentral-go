package definitions

// AccountBusinessAddressResource Account Business Address Resource
type AccountBusinessAddressResource struct {
	Uri string `json:"uri"`
	BusinessAddress ContactBusinessAddressInfo `json:"businessAddress"`
	Company string `json:"company"`
	Email string `json:"email"`
}
