package main

import (
	"github.com/gin-gonic/gin"
	"gogo-boot/go-boot/platform/actuator"
	. "gogo-boot/go-boot/platform/config"
	. "gogo-boot/go-boot/platform/initializer"
	"gogo-boot/go-boot/web/app/authz"
	"gogo-boot/go-boot/web/app/graph"
	_ "gogo-boot/go-boot/web/app/oauth2"
	"gogo-boot/go-boot/web/app/oidc"
	"gogo-boot/go-boot/web/app/openapi"
	"gogo-boot/go-boot/web/app/restapi"
	"gogo-boot/go-boot/web/app/sse"
	"net/http"
)

func main() {

	Router.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	openapi.NewRouter(Router.Group("/openapi"))
	restapi.Routes(Router.Group("/restapi"))
	graph.Routes(Router.Group("/graphql"))
	//oauth2.Routes(Router.Group("/login"))
	oidc.Routes(Router.Group("/login"))
	sse.Routes(Router.Group("/sse"))
	actuator.Routes(Router.Group("/actuator"))
	authz.Routes(Router.Group("/authz"))

	//Todo set Host - only for local test
	Router.Run(AppConfig.Server.Host + ":" + AppConfig.Server.PortNumber)
}
