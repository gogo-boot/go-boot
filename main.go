package main

import (
	"example.com/go-boot/platform/actuator"
	. "example.com/go-boot/platform/config"
	"example.com/go-boot/platform/initializer"
	"example.com/go-boot/web/app/graph"
	"example.com/go-boot/web/app/oidc"
	"example.com/go-boot/web/app/openapi"
	"example.com/go-boot/web/app/restapi"
	"example.com/go-boot/web/app/sse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type load func(*gin.RouterGroup)

// Use package load side effect. When it is loaded, the init function will be initiated
// The init function load "Route" pointer from initializer.
func serviceLoad(fn load) {
}

func main() {

	initializer.Router.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	serviceLoad(openapi.NewRouter)
	serviceLoad(restapi.Routes)
	serviceLoad(graph.Routes)
	//serviceLoad(oauth2.Routes)
	serviceLoad(oidc.Routes)
	serviceLoad(sse.Routes)
	serviceLoad(actuator.Routes)

	//Todo set Host - only for local test
	initializer.Router.Run(":" + AppConfig.Server.PortNumber)
}
