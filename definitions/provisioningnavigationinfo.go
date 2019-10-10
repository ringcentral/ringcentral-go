package definitions

// ProvisioningNavigationInfo Provisioning Navigation Info
type ProvisioningNavigationInfo struct {
	FirstPage ProvisioningNavigationInfoUri `json:"firstPage"`
	NextPage ProvisioningNavigationInfoUri `json:"nextPage"`
	PreviousPage ProvisioningNavigationInfoUri `json:"previousPage"`
	LastPage ProvisioningNavigationInfoUri `json:"lastPage"`
}
