package definitions

// GlipGroupList Glip Group List
type GlipGroupList struct {
	Records []GlipGroupInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
