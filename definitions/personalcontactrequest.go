package definitions

// PersonalContactRequest Personal Contact Request
type PersonalContactRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	MiddleName string `json:"middleName"`
	NickName string `json:"nickName"`
	Company string `json:"company"`
	JobTitle string `json:"jobTitle"`
	Email string `json:"email"`
	Email2 string `json:"email2"`
	Email3 string `json:"email3"`
	Birthday string `json:"birthday"`
	WebPage string `json:"webPage"`
	Notes string `json:"notes"`
	HomePhone string `json:"homePhone"`
	HomePhone2 string `json:"homePhone2"`
	BusinessPhone string `json:"businessPhone"`
	BusinessPhone2 string `json:"businessPhone2"`
	MobilePhone string `json:"mobilePhone"`
	BusinessFax string `json:"businessFax"`
	CompanyPhone string `json:"companyPhone"`
	AssistantPhone string `json:"assistantPhone"`
	CarPhone string `json:"carPhone"`
	OtherPhone string `json:"otherPhone"`
	OtherFax string `json:"otherFax"`
	CallbackPhone string `json:"callbackPhone"`
	HomeAddress ContactAddressInfo `json:"homeAddress"`
	BusinessAddress ContactAddressInfo `json:"businessAddress"`
	OtherAddress ContactAddressInfo `json:"otherAddress"`
}
