package definitions

// SIPInfoResponse SIPInfo Response
type SIPInfoResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
	AuthorizationId string `json:"authorizationId"`
	Domain string `json:"domain"`
	OutboundProxy string `json:"outboundProxy"`
	Transport string `json:"transport"`
	Certificate string `json:"certificate"`
	SwitchBackInterval int `json:"switchBackInterval"`
}
