package definitions

// GlipTeamsList Glip Teams List
type GlipTeamsList struct {
	Records []GlipTeamInfo `json:"records"`
	Navigation GlipNavigationInfo `json:"navigation"`
}
