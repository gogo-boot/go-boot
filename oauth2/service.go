package oauth2

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/endpoints"
	"io"
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

var token *oauth2.Token

var webSSO WebSSO

func Routes(route *gin.Engine) {
	routeGroup := route.Group("/login")
	{
		routeGroup.GET("/", showIndex)
		routeGroup.GET("/login", login)
		routeGroup.GET("/logout", logout)
		routeGroup.GET("/oauth2/code/dbwebsso", loginProcess)
		routeGroup.GET("/info", showTokenInfo)
		routeGroup.GET("/external", getExternalSite)
	}
}

func logout(c *gin.Context) {
	
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

		if webSSO.State != oauthStateString {
			fmt.Errorf("invalid oauth State")
		}
		var err error
		token, err = oauthConfig.Exchange(ctx, webSSO.Code)
		if err != nil {
			fmt.Errorf("code exchange failed: %s", err.Error())
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
