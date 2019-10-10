package definitions

// AuthProfileCheckResource Auth Profile Check Resource
type AuthProfileCheckResource struct {
	Uri string `json:"uri"`
	Successful bool `json:"successful"`
	Details ActivePermissionResource `json:"details"`
}
