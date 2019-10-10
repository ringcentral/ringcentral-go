package definitions

// GlipPersonInfo Glip Person Info
type GlipPersonInfo struct {
	Id string `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	Avatar string `json:"avatar"`
	CompanyId string `json:"companyId"`
	CreationTime string `json:"creationTime"`
	LastModifiedTime string `json:"lastModifiedTime"`
}
