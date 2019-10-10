package definitions

// DeviceModelInfo Device Model Info
type DeviceModelInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Addons []DeviceAddonInfo `json:"addons"`
	Features []string `json:"features"`
}
