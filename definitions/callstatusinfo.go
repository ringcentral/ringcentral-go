package definitions

// CallStatusInfo Call Status Info
type CallStatusInfo struct {
	Code string `json:"code"`
	PeerId PeerInfo `json:"peerId"`
	Reason string `json:"reason"`
	Description string `json:"description"`
}
