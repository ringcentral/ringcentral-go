package definitions

// GetRingOutStatusResponse Get Ring Out Status Response
type GetRingOutStatusResponse struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Status RingOutStatusInfo `json:"status"`
}
