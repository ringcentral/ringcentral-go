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

// RestClient Restful Client
type RestClient struct {
	ClientID     string
	ClientSecret string
	Server       string
	Token        *TokenInfo
}

// Request HTTP request
func (rc RestClient) Request(method, endpoint, contentType string, body io.Reader) []byte {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", rc.Server, endpoint), body)
	if err != nil {
		log.Fatal(err)
	}
	if endpoint == "/restapi/oauth/token" || endpoint == "/restapi/oauth/revoke" {
		req.Header.Add("Authorization", fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", rc.ClientID, rc.ClientSecret)))))
	} else {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", rc.Token.AccessToken))
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
