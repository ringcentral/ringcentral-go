package definitions

// Shipping Shipping
type Shipping struct {
	Address DeviceEmergencyServiceAddressResource `json:"address"`
	Method MethodResource `json:"method"`
	Status string `json:"status"`
	Carrier string `json:"carrier"`
	TrackingNumber string `json:"trackingNumber"`
}
