package definitions

// FavoriteContactList Favorite Contact List
type FavoriteContactList struct {
	Uri string `json:"uri"`
	Records []FavoriteContactResource `json:"records"`
}
