package definitions

// BlockedAllowedPhoneNumbersList Blocked Allowed Phone Numbers List
type BlockedAllowedPhoneNumbersList struct {
	Uri string `json:"uri"`
	Records []BlockedAllowedPhoneNumberInfo `json:"records"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging CallHandlingPagingInfo `json:"paging"`
}
