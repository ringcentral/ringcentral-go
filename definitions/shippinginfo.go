package definitions

// ShippingInfo Shipping Info
type ShippingInfo struct {
	Status string `json:"status"`
	Carrier string `json:"carrier"`
	TrackingNumber string `json:"trackingNumber"`
	Method MethodInfo `json:"method"`
	Address ShippingAddressInfo `json:"address"`
}
