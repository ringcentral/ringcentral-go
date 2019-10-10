package definitions

// ExtensionPermissions Extension Permissions
type ExtensionPermissions struct {
	Admin PermissionInfo `json:"admin"`
	InternationalCalling PermissionInfo `json:"internationalCalling"`
}
