package ringcentral

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

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

// RestClient Restful Client
type RestClient struct {
	ClientID     string
	ClientSecret string
	Server       string
	Token        *TokenInfo
}

// Authorize an user
func (rc *RestClient) Authorize(getTokenRequest GetTokenRequest) {
	body := url.Values{}
	rValue := reflect.ValueOf(&getTokenRequest).Elem()
	rType := rValue.Type()
	for i := 0; i < rValue.NumField(); i++ {
		fieldValue := fmt.Sprint(rValue.Field(i))
		if fieldValue != "" {
			body.Set(rType.Field(i).Tag.Get("json"), fieldValue)
		}
	}
	bytes := rc.Request("POST", "/restapi/oauth/token", "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))
	json.Unmarshal(bytes, &rc.Token)
}

// Revoke revoke current token
func (rc *RestClient) Revoke() {
	if rc.Token == nil {
		return
	}
	rc.Post("/restapi/oauth/revoke", strings.NewReader(fmt.Sprintf(`{token: "%s"}`, rc.Token.AccessToken)))
	rc.Token = nil
}

// Request HTTP request
func (rc RestClient) Request(method, endpoint, contentType string, body io.Reader) []byte {
	req, err := http.NewRequest(method, rc.Server+endpoint, body)
	if err != nil {
		log.Fatal(err)
	}
	// todo: diffrent method to determine token type
	if rc.Token == nil {
		req.Header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(rc.ClientID+":"+rc.ClientSecret)))
	} else {
		req.Header.Add("Authorization", "Bearer "+rc.Token.AccessToken)
	}
	if contentType != "" {
		req.Header.Add("Content-Type", contentType)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}

// Get HTTP GET
func (rc RestClient) Get(endpoint string) []byte {
	return rc.Request("GET", endpoint, "", nil)
}

// Delete HTTP DELETE
func (rc RestClient) Delete(endpoint string) []byte {
	return rc.Request("DELETE", endpoint, "", nil)
}

// Post HTTP POST
func (rc RestClient) Post(endpoint string, body io.Reader) []byte {
	return rc.Request("POST", endpoint, "application/json", body)
}

// Put HTTP PUT
func (rc RestClient) Put(endpoint string, body io.Reader) []byte {
	return rc.Request("PUT", endpoint, "application/json", body)
}

// Patch HTTP PATCH
func (rc RestClient) Patch(endpoint string, body io.Reader) []byte {
	return rc.Request("PATCH", endpoint, "application/json", body)
}
