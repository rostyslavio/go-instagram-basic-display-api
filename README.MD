# Go Instagram Basic Display API
Docs: [Instagram Basic Display API](https://developers.facebook.com/docs/instagram-basic-display-api)
#### Installation
```shell
go get github.com/rostyslavio/go-instagram-basic-display-api
```
## Get Access Tokens and Permissions
### Step 1: Get Authorization
```go
package main

import (
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	redirect, _ := gogram.New().Config(config).GetAuthorizeRedirect()

	log.Printf(redirect)
}
```
### Step 2: Exchange the Code For a Token
```go
package main

import (
	"encoding/json"
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId int `json:"user_id"`
}

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	response, _ := gogram.New().Config(config).GetAccessToken("Query('code')")
	var accessToken AccessToken
	json.Unmarshal([]byte(response), &accessToken)

	log.Printf(response)
}
```

## Get User Profiles and User Media
### Get a User’s Profile
```go
package main

import (
	"encoding/json"
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

type User struct {
	Id string `json:"id"`
	Username int `json:"username"`
}

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	response, _ := gogram.New().Config(config).GetUserProfile("Query('access_token')", []string{"id", "username"})
	var user User
	json.Unmarshal([]byte(response), &user)

	log.Printf(response)
}
```
