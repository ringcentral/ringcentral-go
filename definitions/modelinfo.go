package definitions

// ModelInfo Model Info
type ModelInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Addons []AddonInfo `json:"addons"`
	Features []string `json:"features"`
}
