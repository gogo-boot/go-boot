package main

import (
	. "example.com/go-boot/config"
	"example.com/go-boot/graph"
	"example.com/go-boot/middlewares"
	"example.com/go-boot/oauth2"
	"example.com/go-boot/openapi"
	"example.com/go-boot/restapi"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var server *gin.Engine

func init() {
	logLevel, _ := log.ParseLevel(AppConfig.Server.LogLevel)
	log.SetLevel(logLevel)
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{}) //for easy parsing by logstash or Splunk

	// HTTP Server Set up
	// server = gin.Default() // Default Mode
	server = gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.LoggingMiddleware())

	server.LoadHTMLGlob("template/*")
}

func main() {

	server.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	openapi.NewRouter(server.Group("/openapi"))
	restapi.Routes(server.Group("/restapi"))
	graph.Routes(server.Group("/graphql"))
	oauth2.Routes(server.Group("/login"))
	//myOidc.Routes(server.Group("/login"))
	server.Run(":" + AppConfig.Server.PortNumber)
}
