# Go Instagram Basic Display API
Docs: [Instagram Basic Display API](https://developers.facebook.com/docs/instagram-basic-display-api)
#### Installation
```shell
go get -u github.com/rostyslavio/go-instagram-basic-display-api
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
	redirect, _ := gogram.NewGogram().Config(config).GetAuthorizeRedirect()

	//log.Printf(redirect)
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
	response, err := gogram.NewGogram().Config(config).GetAccessToken("Query('code')")
	if err != nil {
		log.Fatal(err)
	}
	
	var accessToken AccessToken
	json.Unmarshal([]byte(response), &accessToken)
	
	//log.Printf(response)
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
	response, err := gogram.NewGogram().Config(config).GetUserProfile("ACCESS_TOKEN", []string{"id", "username"})
	if err != nil {
		log.Fatal(err)
	}
	
	var user User
	json.Unmarshal([]byte(response), &user)
	
	// log.Printf(response)
}
```
### Get a User’s Media
```go
package main

import (
	"encoding/json"
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

type MediaData struct {
	Id string `json:"id"`
	MediaType string `json:"media_type"`
	MediaUrl string `json:"media_url"`
	ThumbnailUrl *string `json:"thumbnail_url"`
}

type Medias struct {
	Data []MediaData `json:"data"`
}

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	accessToken := "ACCESS_TOKEN"
	response, err := gogram.NewGogram().Config(config).GetUsersMedia(accessToken, []string{"id", "media_type", "media_url", "thumbnail_url"})
	if err != nil {
		log.Fatal(err)
	}
	
	var medias Medias
	json.Unmarshal([]byte(response), &medias)
	
	//s, _ := json.MarshalIndent(medias, "", "  ")
	//log.Printf(string(s))
}
```

### Get a User’s Media (Paging)
```go
package main

import (
	"encoding/json"
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

type MediaData struct {
	Id           string  `json:"id"`
	MediaType    string  `json:"media_type"`
	MediaUrl     string  `json:"media_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
}

type Medias struct {
	Data   []MediaData `json:"data"`
	Paging Paging      `json:"paging"`
}

type Paging struct {
	gogram.Next `json:"next"`
}

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	accessToken := "ACCESS_TOKEN"
	response, err := gogram.NewGogram().Config(config).GetUsersMedia(accessToken, []string{"id", "media_type", "media_url", "thumbnail_url"})
	if err != nil {
		log.Fatal(err)
	}

	var medias Medias
	json.Unmarshal([]byte(response), &medias)

	// Get next page
	response, err = medias.Paging.Next.GetUsersMedia()
	if err != nil {
		log.Fatal(err)
	}
	
	json.Unmarshal([]byte(response), &medias)

	// s, _ := json.MarshalIndent(medias, "", "  ")
	// log.Print(string(s))
}
```
### Get Album Contents
```go
package main

import (
	"encoding/json"
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

type MediaData struct {
	Id string `json:"id"`
	MediaType string `json:"media_type"`
	MediaUrl string `json:"media_url"`
	ThumbnailUrl *string `json:"thumbnail_url"`
}

type Album struct {
	Data []MediaData `json:"data"`
}

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	accessToken := "ACCESS_TOKEN"
	response, err := gogram.NewGogram().Config(config).GetAlbumContents(
		17909521286425942,
		accessToken,
		[]string{"id", "media_type", "media_url", "thumbnail_url"},
	)
	if err != nil {
		log.Fatal(err)
	}

	var album Album
	json.Unmarshal([]byte(response), &album)
	
	//s, _ := json.MarshalIndent(album, "", "  ")
	//log.Printf(string(s))
}
```
## Long-Lived Access Tokens
### Get a Long-Lived Token
```go
package main

import (
	"encoding/json"
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

type LongLivedToken struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	shortLivedAccessToken := "ACCESS_TOKEN"
	response, err := gogram.NewGogram().Config(config).GetLongLivedToken(shortLivedAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	var longLivedToken LongLivedToken
	json.Unmarshal([]byte(response), &longLivedToken)
	
	//s, _ := json.MarshalIndent(longLivedToken, "", "  ")
	//log.Printf(string(s))
}
```
### Refresh a Long-Lived Token
```go
package main

import (
	"encoding/json"
	gogram "github.com/rostyslavio/go-instagram-basic-display-api"
	"log"
)

type LongLivedToken struct {
	AccessToken string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int `json:"expires_in"`
}

func main() {
	config := gogram.Config{
		ClientId: "INSTAGRAM_BASIC_DISPLAY_APP_ID",
		ClientSecret: "INSTAGRAM_BASIC_DISPLAY_APP_SECRET",
		RedirectUri: "INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI",
	}
	longLivedAccessToken := "ACCESS_TOKEN"
	response, err := gogram.NewGogram().Config(config).RefreshLongLivedToken(longLivedAccessToken)
	if err != nil {
		log.Fatal(err)
	}

	var longLivedToken LongLivedToken
	json.Unmarshal([]byte(response), &longLivedToken)
	
	//s, _ := json.MarshalIndent(longLivedToken, "", "  ")
	//log.Printf(string(s))
}
```
## Deauthorize & Data Deletion Requests
### Using a Signed Request
```go
app.Post("/instagram/deauthorize", instagram.Logout())
app.Post("/instagram/deletion", instagram.Logout())
...

type SignedRequest struct {
    SignedRequest string `json:"signed_request" xml:"signed_request" form:"signed_request"`
}

func Logout() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// parse body
		sr := new(SignedRequest)

		if err := c.BodyParser(sr); err != nil {
			return err
		}

		// parse signed request
		config := gogram.Config{
			ClientId: src.Getenv("INSTAGRAM_BASIC_DISPLAY_APP_ID"),
			ClientSecret: src.Getenv("INSTAGRAM_BASIC_DISPLAY_APP_SECRET"),
			RedirectUri: src.Getenv("INSTAGRAM_BASIC_DISPLAY_REDIRECT_URI"),
		}
		data, _ := gogram.NewGogram().Config(config).ParseSignedRequest(sr.SignedRequest)

		// fmt.Println(data)

		return c.SendString("ok")
	}
}
```
