package oidc2

import (
	"context"
	. "example.com/go-boot/config"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"time"
)

var (
	oauthConfig *oauth2.Config
)
var ctx = context.Background()

var verifier *oidc.IDTokenVerifier

func init() {

	provider, err := oidc.NewProvider(ctx, AppConfig.Oidc.Issuer)
	if err != nil {
		log.Panic(err)
	}

	oauthConfig = &oauth2.Config{
		RedirectURL:  AppConfig.Oidc.RedirectUrl,
		ClientID:     AppConfig.Oidc.ClientId,
		ClientSecret: AppConfig.Oidc.ClientSecret,
		Scopes:       AppConfig.Oidc.Scopes,
		Endpoint:     provider.Endpoint(),
	}
	verifier = provider.Verifier(&oidc.Config{ClientID: AppConfig.Oidc.ClientId})
}

var (
	// TODO: randomize it
	oauthStateString = "pseudo-random"
)

type UserInfo struct {
	accessTokenSub        string
	idTokenName           string
	accessTokenExpiration time.Time
	idTokenExpiration     time.Time
	idTokenValue          string
	accessTokenValue      string
}

type WebSSO struct {
	State        string `form:"state"`
	Code         string `form:"code"`
	SessionState string `form:"session_state"`
}

var token *oauth2.Token
var idToken *oidc.IDToken
var webSSO WebSSO

func Routes(rg *gin.RouterGroup) {
	rg.GET("/", showIndex)
	rg.GET("/login", login)
	rg.GET("/logout", logout)
	rg.GET("/oauth2/code/dbwebsso", loginProcess)
	rg.GET("/info", showTokenInfo)
	rg.GET("/external", getExternalSite)
}

func logout(c *gin.Context) {
	c.SetCookie("JSESSIONID", "", 0, "/", "localhost", false, false)
	c.HTML(http.StatusOK, "logout.html", gin.H{})
}

func login(c *gin.Context) {
	url := oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func showIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func loginProcess(c *gin.Context) {

	if c.Bind(&webSSO) == nil {
		var err error
		token, err = oauthConfig.Exchange(ctx, webSSO.Code)
		if err != nil {
			// handle error
		}

		// Extract the ID Token from OAuth2 token.
		rawIDToken, ok := token.Extra("id_token").(string)
		if !ok {
			// handle missing token
		}

		// Parse and verify ID Token payload.
		idToken, err = verifier.Verify(ctx, rawIDToken)
		if err != nil {
			// handle error
		}

		// Extract custom claims
		var claims struct {
			Email    string `json:"email"`
			Verified bool   `json:"email_verified"`
		}
		if err = idToken.Claims(&claims); err != nil {
			// handle error
		}

		response := fmt.Sprintf("<html><body>Login Success and Retriving token is successful<br/></body></html>")
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
	}
}
func showTokenInfo(c *gin.Context) {
	c.HTML(http.StatusOK, "tokeninfo.html", gin.H{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
		"tokenType":    token.TokenType,
		"tokenExpiry":  token.Expiry,
		"idToken":      idToken,
	})
}
func getExternalSite(c *gin.Context) {
	if webSSO.State != oauthStateString {
		fmt.Errorf("invalid oauth State")
	}

	client := oauthConfig.Client(ctx, token)
	response, err := client.Get("https://gateway.hub.db.de/bizhub-api-secured-with-jwt")

	if err != nil {
		fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Errorf("failed reading response body: %s", err.Error())
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", contents)

}
