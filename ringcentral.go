package ringcentral

import "net/http"

// ProductionServer RingCentral production server url
const ProductionServer = "https://platform.ringcentral.com"

// SandboxServer RingCentral sandbox server url
const SandboxServer = "https://platform.devtest.ringcentral.com"

// RestClient RingCentral rest client
type RestClient struct {
	ClientID     string
	ClientSecret string
	Server       string
}

// Get HTTP GET
func (rc RestClient) Get(endpoint string) (*http.Response, error) {
	return http.Get(rc.Server + endpoint)
}
