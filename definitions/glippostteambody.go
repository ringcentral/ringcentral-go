package definitions

// GlipPostTeamBody Glip Post Team Body
type GlipPostTeamBody struct {
	Public bool `json:"public"`
	Name string `json:"name"`
	Description string `json:"description"`
	Members []interface{} `json:"members"`
}
