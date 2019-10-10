package definitions

// GetRingOutStatusResponseIntId Get Ring Out Status Response Int Id
type GetRingOutStatusResponseIntId struct {
	Id int `json:"id"`
	Uri string `json:"uri"`
	Status RingOutStatusInfo `json:"status"`
}
