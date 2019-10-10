package definitions

// CallerIdByDevice Caller Id By Device
type CallerIdByDevice struct {
	Device CallerIdDeviceInfo `json:"device"`
	CallerId CallerIdByDeviceInfo `json:"callerId"`
}
