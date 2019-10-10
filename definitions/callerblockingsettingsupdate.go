package definitions

// CallerBlockingSettingsUpdate Caller Blocking Settings Update
type CallerBlockingSettingsUpdate struct {
	Mode string `json:"mode"`
	NoCallerId string `json:"noCallerId"`
	PayPhones string `json:"payPhones"`
	Greetings []BlockedCallerGreetingInfo `json:"greetings"`
}
