package definitions

// AccountPresenceInfo Account Presence Info
type AccountPresenceInfo struct {
	Uri string `json:"uri"`
	Records []GetPresenceInfo `json:"records"`
	Navigation PresenceNavigationInfo `json:"navigation"`
	Paging PresencePagingInfo `json:"paging"`
}
