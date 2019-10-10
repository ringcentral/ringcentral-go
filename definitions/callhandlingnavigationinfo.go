package definitions

// CallHandlingNavigationInfo Call Handling Navigation Info
type CallHandlingNavigationInfo struct {
	FirstPage CallHandlingNavigationInfoUri `json:"firstPage"`
	NextPage CallHandlingNavigationInfoUri `json:"nextPage"`
	PreviousPage CallHandlingNavigationInfoUri `json:"previousPage"`
	LastPage CallHandlingNavigationInfoUri `json:"lastPage"`
}
