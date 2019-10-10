package definitions

// GetVersionsResponse Get Versions Response
type GetVersionsResponse struct {
	Uri string `json:"uri"`
	ApiVersions []VersionInfo `json:"apiVersions"`
	ServerVersion string `json:"serverVersion"`
	ServerRevision string `json:"serverRevision"`
}
