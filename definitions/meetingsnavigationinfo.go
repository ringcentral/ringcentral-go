package definitions

// MeetingsNavigationInfo Meetings Navigation Info
type MeetingsNavigationInfo struct {
	NextPage MeetingsNavigationInfoUri `json:"nextPage"`
	PreviousPage MeetingsNavigationInfoUri `json:"previousPage"`
	FirstPage MeetingsNavigationInfoUri `json:"firstPage"`
	LastPage MeetingsNavigationInfoUri `json:"lastPage"`
}
