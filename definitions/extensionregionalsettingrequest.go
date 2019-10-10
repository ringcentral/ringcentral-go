package definitions

// ExtensionRegionalSettingRequest Extension Regional Setting Request
type ExtensionRegionalSettingRequest struct {
	HomeCountry ExtensionCountryInfoRequest `json:"homeCountry"`
	Timezone ExtensionTimezoneInfoRequest `json:"timezone"`
	Language ExtensionLanguageInfoRequest `json:"language"`
	GreetingLanguage ExtensionGreetingLanguageInfoRequest `json:"greetingLanguage"`
	FormattingLocale ExtensionFormattingLocaleInfoRequest `json:"formattingLocale"`
	TimeFormat string `json:"timeFormat"`
}
