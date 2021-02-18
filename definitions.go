package ringcentral

// TokenInfo Token Info
type TokenInfo struct {
	AccessToken           string `json:"access_token"`
	ExpiresIn             int    `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int    `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
	TokenType             string `json:"token_type"`
	OwnerID               string `json:"owner_id"`
	EndpointID            string `json:"endpoint_id"`
	IDToken               string `json:"id_token"`
}

// GetTokenRequest Get Token Request
type GetTokenRequest struct {
	Username            string `json:"username"`
	Password            string `json:"password"`
	Extension           string `json:"extension"`
	GrantType           string `json:"grant_type"`
	Code                string `json:"code"`
	RedirectURI         string `json:"redirect_uri"`
	AccessTokenTTL      string `json:"access_token_ttl"`
	RefreshTokenTTL     string `json:"refresh_token_ttl"`
	Scope               string `json:"scope"`
	RefreshToken        string `json:"refresh_token"`
	EndpointID          string `json:"endpoint_id"`
	Pin                 string `json:"pin"`
	ClientID            string `json:"client_id"`
	AccountID           string `json:"account_id"`
	PartnerAccountID    string `json:"partner_account_id"`
	ClientAssertionType string `json:"client_assertion_type"`
	ClientAssertion     string `json:"client_assertion"`
	Assertion           string `json:"assertion"`
	BrandID             string `json:"brand_id"`
	CodeVerifier        string `json:"code_verifier"`
}
