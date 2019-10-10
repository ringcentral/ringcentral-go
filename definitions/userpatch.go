package definitions

// UserPatch User Patch
type UserPatch struct {
	Operations []PatchOperation `json:"Operations"`
	Schemas []string `json:"schemas"`
}
