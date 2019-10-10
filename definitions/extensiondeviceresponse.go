package definitions

// ExtensionDeviceResponse Extension Device Response
type ExtensionDeviceResponse struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Sku string `json:"sku"`
	Type string `json:"type"`
	Name string `json:"name"`
	Status string `json:"status"`
	Serial string `json:"serial"`
	ComputerName string `json:"computerName"`
	Model ModelInfo `json:"model"`
	Extension ExtensionInfoIntId `json:"extension"`
	EmergencyServiceAddress EmergencyServiceAddressResource `json:"emergencyServiceAddress"`
	PhoneLines []PhoneLinesInfo `json:"phoneLines"`
	Shipping ShippingInfo `json:"shipping"`
	BoxBillingId int `json:"boxBillingId"`
	UseAsCommonPhone bool `json:"useAsCommonPhone"`
	LinePooling string `json:"linePooling"`
	InCompanyNet bool `json:"inCompanyNet"`
	Site DeviceSiteInfo `json:"site"`
	LastLocationReportTime string `json:"lastLocationReportTime"`
}
