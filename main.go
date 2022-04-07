package gogram

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	authorizeRedirect = "https://api.instagram.com/oauth/authorize?client_id=%s&redirect_uri=%s&scope=user_profile,user_media&response_type=code"
	accessTokenEndpoint = "https://api.instagram.com/oauth/access_token"
	userProfileEndpoint = "https://graph.instagram.com/me?fields=%s&access_token=%s"
	usersMediaEndpoint = "https://graph.instagram.com/me/media?fields=%s&access_token=%s"
	mediaDataEndpoint = "https://graph.instagram.com/%s?fields=%s&access_token=%s"
	albumContentsEndpoint = "https://graph.instagram.com/%s/children?fields=%s&access_token=%s"
	longLivedTokenEndpoint = "https://graph.instagram.com/access_token?grant_type=ig_exchange_token&client_secret=%s&access_token=%s"
	refreshLongLivedTokenEndpoint = "https://graph.instagram.com/refresh_access_token?grant_type=ig_refresh_token&access_token=%s"
)

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
func (client *GogramClient) GetAuthorizeRedirect() (uri string, err error) {
	return fmt.Sprintf(authorizeRedirect, client.config.ClientId, client.config.RedirectUri), nil
}

// GetAccessToken
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/getting-access-tokens-and-permissions#step-2--exchange-the-code-for-a-token
func (client *GogramClient) GetAccessToken(code string) (response string, err error) {
	data := url.Values{
		"client_id": {client.config.ClientId},
		"client_secret": {client.config.ClientSecret},
		"code": {code},
		"grant_type": {"authorization_code"},
		"redirect_uri": {client.config.RedirectUri},
	}

	res, err := http.PostForm(accessTokenEndpoint, data)

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
func (client *GogramClient) GetUserProfile(accessToken string, fields []string) (response string, err error) {
	f := strings.Join(fields, ",")

	endpoint := fmt.Sprintf(userProfileEndpoint, f, accessToken)

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
func (client *GogramClient) GetUsersMedia(accessToken string, fields []string) (response string, err error) {
	f := strings.Join(fields, ",")

	endpoint := fmt.Sprintf(usersMediaEndpoint, f, accessToken)

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

// Next type for paging
type Next string

// GetUsersMedia (Paging)
func (endpoint Next) GetUsersMedia() (response string, err error) {
	resp, err := http.Get(string(endpoint))

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
func (client *GogramClient) GetMediaData(mediaId int, accessToken string, fields []string) (response string, err error) {
	f := strings.Join(fields, ",")

	endpoint := fmt.Sprintf(mediaDataEndpoint, strconv.Itoa(mediaId), f, accessToken)

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
func (client *GogramClient) GetAlbumContents(mediaId int, accessToken string, fields []string) (response string, err error) {
	f := strings.Join(fields, ",")

	endpoint := fmt.Sprintf(albumContentsEndpoint, strconv.Itoa(mediaId), f, accessToken)

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

// GetLongLivedToken
// https://developers.facebook.com/docs/instagram-basic-display-api/guides/long-lived-access-tokens#get-a-long-lived-token
func (client *GogramClient) GetLongLivedToken(shortLivedAccessToken string) (response string, err error) {
	endpoint := fmt.Sprintf(longLivedTokenEndpoint, client.config.ClientSecret, shortLivedAccessToken)

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
func (client *GogramClient) RefreshLongLivedToken(longLivedAccessToken string) (response string, err error) {
	endpoint := fmt.Sprintf(refreshLongLivedTokenEndpoint, longLivedAccessToken)

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

// ParseSignedRequest
// https://developers.facebook.com/docs/games/gamesonfacebook/login#parsingsr
// https://developers.facebook.com/docs/instagram-basic-display-api/getting-started#deauthorize-callback-url
// https://developers.facebook.com/docs/instagram-basic-display-api/getting-started#data-deletion-request-callback-url
func (client *GogramClient) ParseSignedRequest(sr string) (response string, err error) {
	s := strings.Split(sr, ".")
	encodedSig := s[0]
	encodedData := s[1]

	// decode signature
	sig, err := base64.RawURLEncoding.DecodeString(encodedSig)
	if err != nil {
		return "", err
	}

	// decode data
	data, err := base64.RawURLEncoding.DecodeString(encodedData)
	if err != nil {
		return "", err
	}

	// confirm the signature
	if isValid := ValidMAC([]byte(encodedData), sig, []byte(client.config.ClientSecret)); isValid == false {
		return "", errors.New("Bad signed JSON signature!")
	}

	return string(data), nil
}

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}
