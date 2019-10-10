package definitions

// ContactList Contact List
type ContactList struct {
	Records []PersonalContactResource `json:"records"`
	Navigation UserContactsNavigationInfo `json:"navigation"`
	Paging UserContactsPagingInfo `json:"paging"`
}
