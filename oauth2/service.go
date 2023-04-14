package oauth2

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

var (
	oauthConfig *oauth2.Config
)
var ctx = context.Background()

func init() {
	oauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/login/oauth2/code/dbwebsso",
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Scopes:       []string{"54f98af6-5da1-4d54-8610-8fad122aa628/.default", "openid", "profile", "email"},
		Endpoint:     endpoints.AzureAD("a1a72d9c-49e6-4f6d-9af6-5aafa1183bfd"),
	}
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

func Routes(route *gin.Engine) {
	routegroup := route.Group("/login")
	{
		routegroup.GET("/", func(c *gin.Context) {
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
		})
		routegroup.GET("/login", func(c *gin.Context) {
			url := oauthConfig.AuthCodeURL(oauthStateString, oauth2.AccessTypeOnline)
			c.Redirect(http.StatusTemporaryRedirect, url)
		})
		routegroup.GET("/logout", func(c *gin.Context) {
		})
		routegroup.GET("/oauth2/code/dbwebsso", func(c *gin.Context) {
			var webSSO WebSSO
			if c.Bind(&webSSO) == nil {

				if webSSO.State != oauthStateString {
					fmt.Errorf("invalid oauth State")
				}
				token, err := oauthConfig.Exchange(ctx, webSSO.Code)
				if err != nil {
					fmt.Errorf("code exchange failed: %s", err.Error())
				}

				response := fmt.Sprintf("<html><body>accesstoken : %s<br/>"+
					"refreshtoken : %s<br/>"+
					"tokentype: %s<br/>"+
					"tokenexpiry : %s<br/></body></html>", token.AccessToken, token.RefreshToken, token.TokenType, token.Expiry)
				c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(response))
				//c.Data(http.StatusOK, "text/html; charset=utf-8", string(response))
			}
		})
		routegroup.GET("/info", func(c *gin.Context) {

		})
		routegroup.GET("/external", getExternalSite)
	}
}

func getExternalSite(c *gin.Context) {
	var webSSO WebSSO
	if c.ShouldBind(&webSSO) == nil {
		if webSSO.State != oauthStateString {
			fmt.Errorf("invalid oauth State")
		}
		token, err := oauthConfig.Exchange(ctx, webSSO.Code)
		if err != nil {
			fmt.Errorf("Code exchange failed: %s", err.Error())
		}
		client := oauthConfig.Client(ctx, token)
		response, err := client.Get("https://gateway.hub.db.de/bizhub-api-secured-with-jwt")

		if err != nil {
			fmt.Errorf("failed getting user info: %s", err.Error())
		}
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			fmt.Errorf("failed reading response body: %s", err.Error())
		}
		c.HTML(http.StatusOK, "text/html; charset=utf-8", contents)
	}
}
