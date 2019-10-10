package definitions

// AuthProfileResource Auth Profile Resource
type AuthProfileResource struct {
	Uri string `json:"uri"`
	Permissions []ActivePermissionResource `json:"permissions"`
}
