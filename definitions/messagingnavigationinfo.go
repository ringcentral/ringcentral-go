package definitions

// MessagingNavigationInfo Messaging Navigation Info
type MessagingNavigationInfo struct {
	FirstPage MessagingNavigationInfoURI `json:"firstPage"`
	NextPage MessagingNavigationInfoURI `json:"nextPage"`
	PreviousPage MessagingNavigationInfoURI `json:"previousPage"`
	LastPage MessagingNavigationInfoURI `json:"lastPage"`
}
