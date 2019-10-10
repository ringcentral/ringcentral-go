package definitions

// BrandInfo Brand Info
type BrandInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	HomeCountry CountryInfo `json:"homeCountry"`
}
