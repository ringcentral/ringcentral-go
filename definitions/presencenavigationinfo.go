package definitions

// PresenceNavigationInfo Presence Navigation Info
type PresenceNavigationInfo struct {
	FirstPage PresenceNavigationInfoURI `json:"firstPage"`
	NextPage PresenceNavigationInfoURI `json:"nextPage"`
	PreviousPage PresenceNavigationInfoURI `json:"previousPage"`
	LastPage PresenceNavigationInfoURI `json:"lastPage"`
}
