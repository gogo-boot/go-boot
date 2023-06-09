package oidc

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gogo-boot/go-boot/platform/authenticator"
	. "gogo-boot/go-boot/platform/config"
	"gogo-boot/go-boot/platform/middleware"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
)

func Routes(rg *gin.RouterGroup) {

	auth, err := authenticator.NewOidc()
	if err != nil {
		log.Fatalf("Failed to initialize the authenticator: %v", err)
	}

	rg.GET("/", showIndex)
	rg.GET("/loginHome", showHome)
	rg.GET("/login", loginHandler(auth))
	rg.GET("/oauth2/code/dbwebsso", callBackHandler(auth))
	rg.GET("/user", middleware.IsAuthenticated, showUserInfo)
	rg.GET("/logout", logoutHandler)
	rg.GET("/info", showTokenInfo)
	rg.GET("/external", getExternalSite(auth))
}

func logoutHandler(ctx *gin.Context) {

	logoutUrl, err := url.Parse("https://login.microsoftonline.com/a1a72d9c-49e6-4f6d-9af6-5aafa1183bfd/oauth2/v2.0/logout")
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

	session := sessions.Default(ctx)
	session.Clear()

	ctx.Redirect(http.StatusTemporaryRedirect, logoutUrl.String())
}

// Handler for our login.
func loginHandler(auth *authenticator.Authenticator) gin.HandlerFunc {
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
	c.HTML(http.StatusOK, "login.html", nil)
}

func showHome(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

// Handler for our callback.
func callBackHandler(auth *authenticator.Authenticator) gin.HandlerFunc {
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

		var profile struct {
			FamilyName string   `json:"family_name"`
			GivenName  string   `json:"given_name"`
			Groups     []string `json:"groups"`
			Email      string   `json:"email"`
			Name       string   `json:"name"`
			Roles      []string `json:"roles"`
		}
		if err := idToken.Claims(&profile); err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}

		session.Set("token", token)
		session.Set("profile", profile)
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

	// convert struct to map
	var myMap map[string]interface{}
	data, _ := json.Marshal(profile)
	json.Unmarshal(data, &myMap)

	ctx.HTML(http.StatusOK, "user.html", myMap)
}

func showTokenInfo(ctx *gin.Context) {
	session := sessions.Default(ctx)
	sessionToken := session.Get("token")
	profile := session.Get("profile")
	// convert struct to map
	var myMap map[string]interface{}
	data, _ := json.Marshal(profile)
	json.Unmarshal(data, &myMap)

	token := sessionToken.(oauth2.Token)

	ctx.HTML(http.StatusOK, "tokeninfo.html", gin.H{
		"accessToken":  token.AccessToken,
		"refreshToken": token.RefreshToken,
		"tokenType":    token.TokenType,
		"tokenExpiry":  token.Expiry,
		"userProfile":  profile,
	})
}

func getExternalSite(auth *authenticator.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionToken := session.Get("token")
		token := sessionToken.(oauth2.Token)

		client := auth.Config.Client(ctx, &token)
		response, err := client.Get("https://gateway.hub.db.de/bizhub-api-secured-with-jwt")

		if err != nil {
			fmt.Errorf("failed getting user info: %s", err.Error())
		}
		defer response.Body.Close()
		contents, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Errorf("failed reading response body: %s", err.Error())
		}
		ctx.Data(http.StatusOK, "text/html; charset=utf-8", contents)
	}
}
