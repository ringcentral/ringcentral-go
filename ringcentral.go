package ringcentral

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ringcentral/ringcentral-go/definitions"
)

// ProductionServer RingCentral production server url
const ProductionServer = "https://platform.ringcentral.com"

// SandboxServer RingCentral sandbox server url
const SandboxServer = "https://platform.devtest.ringcentral.com"

// RestClient RingCentral Restful client
type RestClient struct {
	ClientID     string
	ClientSecret string
	Server       string
	Token        *definitions.TokenInfo
}

// Authorize an user
func (rc *RestClient) Authorize(username string, extension string, password string) {
	data := url.Values{
		"grant_type": []string{"password"},
		"username":   []string{username},
		"extension":  []string{extension},
		"password":   []string{password},
	}
	bytes := rc.Request("POST", "/restapi/oauth/token", "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
	json.Unmarshal(bytes, &rc.Token)
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

// Request HTTP request
func (rc RestClient) Request(method, endpoint, contentType string, body io.Reader) []byte {
	req, err := http.NewRequest(method, rc.Server+endpoint, body)
	if err != nil {
		log.Fatal(err)
	}
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
