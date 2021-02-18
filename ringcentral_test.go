package ringcentral

import (
	"os"
	"testing"
)

func TestAuthorize(t *testing.T) {
	rc := RestClient{
		ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
		ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
		Server:       os.Getenv("RINGCENTRAL_SERVER_URL"),
	}

	rc.Authorize(GetTokenRequest{
		GrantType: "password",
		Username:  os.Getenv("RINGCENTRAL_USERNAME"),
		Extension: os.Getenv("RINGCENTRAL_EXTENSION"),
		Password:  os.Getenv("RINGCENTRAL_PASSWORD"),
	})

	if len(rc.Token.AccessToken) <= 0 {
		t.Error("access token length invalid")
	}

	rc.Revoke()

	if rc.Token != nil {
		t.Error("Revoke failed")
	}
}
