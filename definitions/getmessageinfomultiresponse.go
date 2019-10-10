package definitions

// GetMessageInfoMultiResponse Get Message Info Multi Response
type GetMessageInfoMultiResponse struct {
	ResourceId string `json:"resourceId"`
	Status int `json:"status"`
	Body MessageBody `json:"body"`
}
