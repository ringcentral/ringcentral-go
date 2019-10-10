package definitions

// PersonalContactResource Personal Contact Resource
type PersonalContactResource struct {
	Uri string `json:"uri"`
	Availability string `json:"availability"`
	Email string `json:"email"`
	Id int `json:"id"`
	Notes string `json:"notes"`
	Company string `json:"company"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	JobTitle string `json:"jobTitle"`
	Birthday string `json:"birthday"`
	WebPage string `json:"webPage"`
	MiddleName string `json:"middleName"`
	NickName string `json:"nickName"`
	Email2 string `json:"email2"`
	Email3 string `json:"email3"`
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
	BusinessAddress ContactAddressInfo `json:"businessAddress"`
	HomeAddress ContactAddressInfo `json:"homeAddress"`
	OtherAddress ContactAddressInfo `json:"otherAddress"`
}
