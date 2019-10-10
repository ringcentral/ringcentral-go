package definitions

// ServiceProviderConfig Service Provider Config
type ServiceProviderConfig struct {
	AuthenticationSchemes []AuthenticationScheme `json:"authenticationSchemes"`
	Bulk BulkSupported `json:"bulk"`
	ChangePassword Supported `json:"changePassword"`
	Etag Supported `json:"etag"`
	Filter FilterSupported `json:"filter"`
	Patch Supported `json:"patch"`
	Schemas []string `json:"schemas"`
	Sort Supported `json:"sort"`
	XmlDataFormat Supported `json:"xmlDataFormat"`
}
