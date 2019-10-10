package definitions

// ProfileImageResource Profile Image Resource
type ProfileImageResource struct {
	Uri string `json:"uri"`
	Etag string `json:"etag"`
	ContentType string `json:"contentType"`
	LastModified string `json:"lastModified"`
	Scales []ScaledProfileImageResource `json:"scales"`
}
