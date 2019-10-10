package definitions

// RegionalSettings Regional Settings
type RegionalSettings struct {
	HomeCountry CountryInfo `json:"homeCountry"`
	Timezone TimezoneInfo `json:"timezone"`
	Language LanguageInfo `json:"language"`
	GreetingLanguage GreetingLanguageInfo `json:"greetingLanguage"`
	FormattingLocale FormattingLocaleInfo `json:"formattingLocale"`
	TimeFormat string `json:"timeFormat"`
}
