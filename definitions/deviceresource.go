package definitions

// DeviceResource Device Resource
type DeviceResource struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Sku string `json:"sku"`
	Type string `json:"type"`
	Status string `json:"status"`
	Name string `json:"name"`
	Serial string `json:"serial"`
	ComputerName string `json:"computerName"`
	Model ModelInfo `json:"model"`
	Extension ExtensionResourceIntId `json:"extension"`
	PhoneLines []PhoneLineResource `json:"phoneLines"`
	EmergencyServiceAddress EmergencyServiceAddressResource `json:"emergencyServiceAddress"`
	Shipping ShippingResource `json:"shipping"`
	BoxBillingId int `json:"boxBillingId"`
	LinePooling string `json:"linePooling"`
	UseAsCommonPhone bool `json:"useAsCommonPhone"`
}
