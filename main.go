package main

import (
	"example.com/go-boot/actuator"
	. "example.com/go-boot/config"
	"example.com/go-boot/graph"
	"example.com/go-boot/initializer"
	"example.com/go-boot/openapi"
	"example.com/go-boot/restapi"
	"example.com/go-boot/sse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type load func(*gin.RouterGroup)

func serviceLoad(fn load) {
}

func main() {

	serviceLoad(openapi.NewRouter)
	serviceLoad(restapi.Routes)
	serviceLoad(graph.Routes)
	//service_load(oauth2.Routes)
	//service_load(oidc.Routes)
	serviceLoad(sse.Routes)
	serviceLoad(actuator.Routes)
	initializer.Router.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	//graph.Routes(Router.Group("/graphql"))
	////oauth2.Routes(Router.Group("/login"))
	//oidc.Routes(Router.Group("/login"))
	//Todo set Host - only for local test
	initializer.Router.Run(":" + AppConfig.Server.PortNumber)
}
