package definitions

// BulkSupported Bulk Supported
type BulkSupported struct {
	MaxOperations int `json:"maxOperations"`
	MaxPayloadSize int `json:"maxPayloadSize"`
	Supported bool `json:"supported"`
}
