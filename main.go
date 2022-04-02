package gogram

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Config struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
}

// GogramClient is the main struct of the package
type GogramClient struct {
	config Config
}

func NewGogram() *GogramClient {
	return &GogramClient{}
}

// Config is needed to success api work
func (client *GogramClient) Config(config Config) *GogramClient {
	client.config = config
	return client
}

// GetAuthorizeRedirect
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-access-tokens-and-permissions#step-1--get-authorization
func (client *GogramClient) GetAuthorizeRedirect() (string, error) {
	return "https://api.instagram.com/oauth/authorize" +
		"?client_id=" + client.config.ClientId +
		"&redirect_uri=" + client.config.RedirectUri +
		"&scope=user_profile,user_media" +
		"&response_type=code", nil
}

// GetAccessToken
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-access-tokens-and-permissions#step-2--exchange-the-code-for-a-token
func (client *GogramClient) GetAccessToken(code string) (string, error) {
	data := url.Values{
		"client_id": {client.config.ClientId},
		"client_secret": {client.config.ClientSecret},
		"code": {code},
		"grant_type": {"authorization_code"},
		"redirect_uri": {client.config.RedirectUri},
	}

	res, err := http.PostForm("https://api.instagram.com/oauth/access_token", data)

	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetUserProfile
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-profiles-and-media#get-a-user-s-profile
// https://developers.facebook.com/docs/instagram-basic-display-api/reference/user#fields
func (client *GogramClient) GetUserProfile(accessToken string, fields []string) (string, error) {
	f := strings.Join(fields, ",")

	endpoint := "https://graph.instagram.com/me" +
		"?fields=" + f +
		"&access_token=" + accessToken

	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetUsersMedia
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-profiles-and-media#get-a-user-s-media
// https://developers.facebook.com/docs/instagram-basic-display-api/reference/media#fields
func (client *GogramClient) GetUsersMedia(accessToken string, fields []string, nextPage ...string) (string, error) {
	f := strings.Join(fields, ",")

	endpoint := "https://graph.instagram.com/me/media" +
		"?fields=" + f +
		"&access_token=" + accessToken

	if len(nextPage) > 0 {
		endpoint = nextPage[0]
	}

	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetMediaData
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-profiles-and-media#get-media-data
// https://developers.facebook.com/docs/instagram-basic-display-api/reference/media#fields
func (client *GogramClient) GetMediaData(mediaId int, accessToken string, fields []string) (string, error) {
	f := strings.Join(fields, ",")

	endpoint := "https://graph.instagram.com/" + strconv.Itoa(mediaId) +
		"?fields=" + f +
		"&access_token=" + accessToken

	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetAlbumContents
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-profiles-and-media#get-album-contents
// https://developers.facebook.com/docs/instagram-basic-display-api/reference/media#fields
func (client *GogramClient) GetAlbumContents(mediaId int, accessToken string, fields []string) (string, error) {
	f := strings.Join(fields, ",")

	endpoint := "https://graph.instagram.com/" + strconv.Itoa(mediaId) + "/children" +
		"?fields=" + f +
		"&access_token=" + accessToken

	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// GetLongLovedToken
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/long-lived-access-tokens#get-a-long-lived-token
func (client *GogramClient) GetLongLovedToken(shortLivedAccessToken string) (string, error) {
	endpoint := "https://graph.instagram.com/access_token" +
		"?grant_type=ig_exchange_token" +
		"&client_secret=" + (client.config.ClientSecret) +
		"&access_token=" + shortLivedAccessToken

	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}

// RefreshLongLivedToken
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/long-lived-access-tokens#refresh-a-long-lived-token
func (client *GogramClient) RefreshLongLivedToken(longLivedAccessToken string) (string, error) {
	endpoint := "https://graph.instagram.com/refresh_access_token" +
		"?grant_type=ig_refresh_token" +
		"&access_token=" + longLivedAccessToken

	resp, err := http.Get(endpoint)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(body), nil
}
