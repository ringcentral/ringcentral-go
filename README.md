# RingCentral GoLang SDK

Experimental, not ready for production!


## Usage

### GET

```go
package main

import (
	"os"

	"github.com/ringcentral/ringcentral-go"
)

func main() {
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
	bytes := rc.Get("/restapi/v1.0/account/~/extension/~")
	println(string(bytes))
}
```

### POST

```go
bytes := rc.Post("/restapi/v1.0/client-info/sip-provision", strings.NewReader(`{"sipInfo":[{"transport":"WSS"}]}`))
println(string(bytes))
println(string(bytes))
```
