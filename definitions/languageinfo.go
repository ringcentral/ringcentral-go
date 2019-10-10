package definitions

// LanguageInfo Language Info
type LanguageInfo struct {
	Id string `json:"id"`
	Uri string `json:"uri"`
	Greeting bool `json:"greeting"`
	FormattingLocale bool `json:"formattingLocale"`
	LocaleCode string `json:"localeCode"`
	Name string `json:"name"`
	Ui bool `json:"ui"`
}
