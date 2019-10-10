package definitions

// RegionalSettingsInfo Regional Settings Info
type RegionalSettingsInfo struct {
	Timezone TimezoneResource `json:"timezone"`
	HomeCountry CountryResource `json:"homeCountry"`
	Language LanguageResource `json:"language"`
	GreetingLanguage LanguageResource `json:"greetingLanguage"`
	FormattingLocale LanguageResource `json:"formattingLocale"`
}
