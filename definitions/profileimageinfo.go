package definitions

// ProfileImageInfo Profile Image Info
type ProfileImageInfo struct {
	Uri string `json:"uri"`
	Etag string `json:"etag"`
	LastModified string `json:"lastModified"`
	ContentType string `json:"contentType"`
	Scales []ProfileImageInfoURI `json:"scales"`
}
