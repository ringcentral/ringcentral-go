package definitions

// ShippingResource Shipping Resource
type ShippingResource struct {
	Address EmergencyServiceAddressResource `json:"address"`
	Method MethodResource `json:"method"`
	Status string `json:"status"`
	Carrier string `json:"carrier"`
	TrackingNumber string `json:"trackingNumber"`
}
