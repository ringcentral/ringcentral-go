package definitions

// GetVersionResponse Get Version Response
type GetVersionResponse struct {
	Uri string `json:"uri"`
	VersionString string `json:"versionString"`
	ReleaseDate string `json:"releaseDate"`
	UriString string `json:"uriString"`
}
