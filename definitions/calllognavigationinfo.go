package definitions

// CallLogNavigationInfo Call Log Navigation Info
type CallLogNavigationInfo struct {
	FirstPage CallLogNavigationInfoURI `json:"firstPage"`
	NextPage CallLogNavigationInfoURI `json:"nextPage"`
	PreviousPage CallLogNavigationInfoURI `json:"previousPage"`
	LastPage CallLogNavigationInfoURI `json:"lastPage"`
}
