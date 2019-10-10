package definitions

// TokenInfo Token Info
type TokenInfo struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	TokenType             string `json:"token_type"`
	OwnerId               string `json:"owner_id"`
	EndpointId            string `json:"endpoint_id"`
	IdToken               string `json:"id_token"`
}
