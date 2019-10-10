package definitions

// UserSearchResponse User Search Response
type UserSearchResponse struct {
	Resources []UserResponse `json:"Resources"`
	ItemsPerPage int `json:"itemsPerPage"`
	Schemas []string `json:"schemas"`
	StartIndex int `json:"startIndex"`
	TotalResults int `json:"totalResults"`
}
