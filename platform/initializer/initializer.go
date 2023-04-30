package initializer

import (
	. "example.com/go-boot/platform/config"
	"example.com/go-boot/platform/middleware"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var Router *gin.Engine

func init() {

	logLevel, _ := log.ParseLevel(AppConfig.Server.LogLevel)
	log.SetLevel(logLevel)
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{}) //for easy parsing by logstash or Splunk

	// HTTP Server Set up
	// Router = gin.Default() // Default Mode
	Router = gin.New()
	Router.Use(gin.Recovery())
	Router.Use(middleware.LoggingMiddleware())

	// Load HTML Template
	Router.LoadHTMLGlob("web/template/*")
}
