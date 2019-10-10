package definitions

// PhoneNumberInfoNumberParser Phone Number Info Number Parser
type PhoneNumberInfoNumberParser struct {
	AreaCode string `json:"areaCode"`
	Country GetCountryInfoNumberParser `json:"country"`
	Dialable string `json:"dialable"`
	E164 string `json:"e164"`
	FormattedInternational string `json:"formattedInternational"`
	FormattedNational string `json:"formattedNational"`
	OriginalString string `json:"originalString"`
	Special bool `json:"special"`
	Normalized string `json:"normalized"`
	TollFree bool `json:"tollFree"`
	SubAddress string `json:"subAddress"`
	SubscriberNumber string `json:"subscriberNumber"`
	DtmfPostfix string `json:"dtmfPostfix"`
}
