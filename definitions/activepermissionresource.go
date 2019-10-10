package definitions

// ActivePermissionResource Active Permission Resource
type ActivePermissionResource struct {
	Permission PermissionIdResource `json:"permission"`
	EffectiveRole RoleIdResource `json:"effectiveRole"`
	Scopes []string `json:"scopes"`
}
