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

var router *gin.Engine

func init() {
	logLevel, _ := log.ParseLevel(AppConfig.Server.LogLevel)
	log.SetLevel(logLevel)
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{}) //for easy parsing by logstash or Splunk

	// HTTP Server Set up
	// router = gin.Default() // Default Mode
	router = gin.New()
	router.Use(gin.Recovery())
	router.Use(middlewares.LoggingMiddleware())

	// Load HTML Template
	router.LoadHTMLGlob("template/*")
}

func main() {

	router.Any("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	openapi.NewRouter(router.Group("/openapi"))
	restapi.Routes(router.Group("/restapi"))
	graph.Routes(router.Group("/graphql"))
	oauth2.Routes(router.Group("/login"))
	//myOidc.Routes(router.Group("/login"))
	router.Run("127.0.0.1:" + AppConfig.Server.PortNumber)
}
