package definitions

// CallerBlockingSettings Caller Blocking Settings
type CallerBlockingSettings struct {
	Mode string `json:"mode"`
	NoCallerId string `json:"noCallerId"`
	PayPhones string `json:"payPhones"`
	Greetings []BlockedCallerGreetingInfo `json:"greetings"`
}
