package definitions

// ExtensionPermissionsResource Extension Permissions Resource
type ExtensionPermissionsResource struct {
	Uri string `json:"uri"`
	Admin Permission `json:"admin"`
	InternationalCalling Permission `json:"internationalCalling"`
	FreeSoftPhoneDigitalLine Permission `json:"freeSoftPhoneDigitalLine"`
}
