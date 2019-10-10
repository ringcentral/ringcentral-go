package definitions

// ScimErrorResponse Scim Error Response
type ScimErrorResponse struct {
	Detail string `json:"detail"`
	Schemas []string `json:"schemas"`
	ScimType string `json:"scimType"`
	Status string `json:"status"`
}
