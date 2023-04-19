package oidc2

import (
	"context"
	. "example.com/go-web-template/config"
	"fmt"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
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
		// handle error
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
	var htmlIndex = []byte(`<html>
				<body>
						Logout Success<br/>
				</body>
				</html>`)
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlIndex)
}

func login(c *gin.Context) {
	url := oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func showIndex(c *gin.Context) {
	var htmlIndex = []byte(`<html>
				<body>
						<br><a href="/login/login">login</a>
						<br><a href="/login/logout">logout</a>
						<br><a href="/login/info">show token info</a>
						<br><a href="/login/external">call external service</a>
						<br><a href="/login/only-with-role">only with role</a>
				</body>
				</html>`)
	c.Data(http.StatusOK, "text/html; charset=utf-8", htmlIndex)
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

	response := fmt.Sprintf("<html><body>accesstoken : %s<br/>"+
		"refreshtoken : %s<br/>"+
		"tokentype: %s<br/>"+
		"tokenexpiry : %s<br/></body></html>", token.AccessToken, token.RefreshToken, token.TokenType, token.Expiry)
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
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
