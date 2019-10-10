package definitions

// UserContactsNavigationInfo User Contacts Navigation Info
type UserContactsNavigationInfo struct {
	FirstPage UserContactsNavigationInfoUri `json:"firstPage"`
	NextPage UserContactsNavigationInfoUri `json:"nextPage"`
	PreviousPage UserContactsNavigationInfoUri `json:"previousPage"`
	LastPage UserContactsNavigationInfoUri `json:"lastPage"`
}
