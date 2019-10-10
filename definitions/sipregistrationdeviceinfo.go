package definitions

// SipRegistrationDeviceInfo Sip Registration Device Info
type SipRegistrationDeviceInfo struct {
	Uri string `json:"uri"`
	Id string `json:"id"`
	Type string `json:"type"`
	Sku string `json:"sku"`
	Status string `json:"status"`
	Name string `json:"name"`
	Serial string `json:"serial"`
	ComputerName string `json:"computerName"`
	Model DeviceModelInfo `json:"model"`
	Extension DeviceExtensionInfo `json:"extension"`
	EmergencyServiceAddress DeviceEmergencyServiceAddressResource `json:"emergencyServiceAddress"`
	Shipping Shipping `json:"shipping"`
	PhoneLines []DevicePhoneLinesInfo `json:"phoneLines"`
	BoxBillingId int `json:"boxBillingId"`
	UseAsCommonPhone bool `json:"useAsCommonPhone"`
	LinePooling string `json:"linePooling"`
	InCompanyNet bool `json:"inCompanyNet"`
	Site DeviceSiteInfo `json:"site"`
	LastLocationReportTime string `json:"lastLocationReportTime"`
}
