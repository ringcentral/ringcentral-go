package definitions

// CallQueues Call Queues
type CallQueues struct {
	Uri string `json:"uri"`
	Records []CallQueueInfo `json:"records"`
	Navigation ProvisioningNavigationInfo `json:"navigation"`
	Paging ProvisioningPagingInfo `json:"paging"`
}
