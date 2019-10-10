package definitions

// GetExtensionForwardingNumberListResponse Get Extension Forwarding Number List Response
type GetExtensionForwardingNumberListResponse struct {
	Records []ForwardingNumberInfo `json:"records"`
	Navigation CallHandlingNavigationInfo `json:"navigation"`
	Paging CallHandlingPagingInfo `json:"paging"`
}
