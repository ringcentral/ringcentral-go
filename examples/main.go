package main

import (
	"encoding/json"
	"log"
	"os"
	"os/user"
	"strings"

	"github.com/joho/godotenv"
	"github.com/ringcentral/ringcentral"
	"github.com/ringcentral/ringcentral/definitions"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Overload(usr.HomeDir + "/.env.prod")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	rc := ringcentral.RestClient{
		ClientID:     os.Getenv("RINGCENTRAL_CLIENT_ID"),
		ClientSecret: os.Getenv("RINGCENTRAL_CLIENT_SECRET"),
		Server:       os.Getenv("RINGCENTRAL_SERVER_URL"),
	}
	rc.Authorize(
		os.Getenv("RINGCENTRAL_USERNAME"),
		os.Getenv("RINGCENTRAL_EXTENSION"),
		os.Getenv("RINGCENTRAL_PASSWORD"),
	)

	bytes := rc.Post("/restapi/v1.0/client-info/sip-provision", strings.NewReader(`{"sipInfo":[{"transport":"WSS"}]}`))
	println(string(bytes))
	var createSipRegistrationResponse definitions.CreateSipRegistrationResponse
	json.Unmarshal(bytes, &createSipRegistrationResponse)
	println(createSipRegistrationResponse.SipInfo[0].Username)
}
