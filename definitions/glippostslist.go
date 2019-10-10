package definitions

// GlipPostsList Glip Posts List
type GlipPostsList struct {
	Records []GlipPostInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
