package definitions

// PatchOperation Patch Operation
type PatchOperation struct {
	Op string `json:"op"`
	Path string `json:"path"`
	Value string `json:"value"`
}
