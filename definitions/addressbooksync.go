package definitions

// AddressBookSync Address Book Sync
type AddressBookSync struct {
	Uri string `json:"uri"`
	Records []PersonalContactResource `json:"records"`
	SyncInfo SyncInfo `json:"syncInfo"`
	NextPageId int `json:"nextPageId"`
	NextPageUri string `json:"nextPageUri"`
}
