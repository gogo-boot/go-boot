package oidc

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"errors"
	. "example.com/go-boot/platform/config"
	"example.com/go-boot/platform/initializer"
	"example.com/go-boot/platform/middleware"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Authenticator is used to authenticate our users.
type Authenticator struct {
	*oidc.Provider
	oauth2.Config
}

// New instantiates the *Authenticator.
func New() (*Authenticator, error) {
	provider, err := oidc.NewProvider(
		context.Background(),
		AppConfig.Oidc.Issuer,
	)
	if err != nil {
		return nil, err
	}

	conf := oauth2.Config{
		RedirectURL:  AppConfig.Oidc.RedirectUrl,
		ClientID:     AppConfig.Oidc.ClientId,
		ClientSecret: AppConfig.Oidc.ClientSecret,
		Endpoint:     provider.Endpoint(),
		Scopes:       AppConfig.Oidc.Scopes,
	}

	return &Authenticator{
		Provider: provider,
		Config:   conf,
	}, nil
}

// VerifyIDToken verifies that an *oauth2.Token is a valid *oidc.IDToken.
func (a *Authenticator) VerifyIDToken(ctx context.Context, token *oauth2.Token) (*oidc.IDToken, error) {
	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, errors.New("no id_token field in oauth2 token")
	}

	oidcConfig := &oidc.Config{
		ClientID: a.ClientID,
	}

	return a.Verifier(oidcConfig).Verify(ctx, rawIDToken)
}

func init() {
	Routes(initializer.Router.Group("/login"))
}

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
var rawIDToken string

func Routes(rg *gin.RouterGroup) {

	auth, err := New()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rg.GET("/", showIndex)
	rg.GET("/login", loginHandler(auth))
	rg.GET("/oauth2/code/dbwebsso", callBackHandler(auth))
	rg.GET("/user", middleware.IsAuthenticated, showUserInfo)
	rg.GET("/logout", logoutHandler)
	rg.GET("/info", showTokenInfo)
	//rg.GET("/external", getExternalSite)
}

func logoutHandler(ctx *gin.Context) {

	logoutUrl, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/v2/logout")
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + ctx.Request.Host)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
		return
	}

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", AppConfig.Oidc.ClientId)
	logoutUrl.RawQuery = parameters.Encode()

	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}

// Handler for our login.
func loginHandler(auth *Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state, err := generateRandomState()
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Save the state inside the session.
		session := sessions.Default(ctx)
		session.Set("state", state)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Redirect(http.StatusTemporaryRedirect, auth.AuthCodeURL(state))
	}
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}

func showIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

// Handler for our callback.
func callBackHandler(auth *Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if ctx.Query("state") != session.Get("state") {
			ctx.String(http.StatusBadRequest, "Invalid state parameter.")
			return
		}

		// Exchange an authorization code for a token.
		token, err := auth.Exchange(ctx.Request.Context(), ctx.Query("code"))
		if err != nil {
			ctx.String(http.StatusUnauthorized, "Failed to convert an authorization code into a token.")
			return
		}

		idToken, err := auth.VerifyIDToken(ctx.Request.Context(), token)
		if err != nil {
			ctx.String(http.StatusInternalServerError, "Failed to verify ID Token.")
			return
		}

		var profile map[string]interface{}
		if err := idToken.Claims(&profile); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		session.Set("access_token", token.AccessToken)
		//session.Set("profile", profile)
		if err := session.Save(); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		// Redirect to logged in page.
		ctx.Redirect(http.StatusTemporaryRedirect, "/login/user")
	}
}

// Handler for our logged-in user page.
func showUserInfo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	profile := session.Get("profile")

	ctx.HTML(http.StatusOK, "user.html", profile)
}

func showTokenInfo(c *gin.Context) {

	c.HTML(http.StatusOK, "tokeninfo.html", gin.H{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
		"tokenType":    token.TokenType,
		"tokenExpiry":  token.Expiry,
		"idToken":      rawIDToken,
	})
}

/*func getExternalSite(c *gin.Context) {
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
*/
