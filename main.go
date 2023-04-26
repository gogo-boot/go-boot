package main

import (
	"example.com/go-boot/actuator"
	. "example.com/go-boot/config"
	"example.com/go-boot/graph"
	"example.com/go-boot/initializer"
	"example.com/go-boot/oidc"
	"example.com/go-boot/openapi"
	"example.com/go-boot/restapi"
	"example.com/go-boot/sse"
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
