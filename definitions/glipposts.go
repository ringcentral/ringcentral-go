package definitions

// GlipPosts Glip Posts
type GlipPosts struct {
	Records []GlipPostInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
