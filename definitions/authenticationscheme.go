package definitions

// AuthenticationScheme Authentication Scheme
type AuthenticationScheme struct {
	Description string `json:"description"`
	DocumentationUri string `json:"documentationUri"`
	Name string `json:"name"`
	SpecUri string `json:"specUri"`
	Primary bool `json:"primary"`
}
