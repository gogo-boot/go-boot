package main

import (
	. "example.com/go-web-template/config"
	"example.com/go-web-template/gorm"
	"example.com/go-web-template/graphql"
	"example.com/go-web-template/middlewares"
	"example.com/go-web-template/oauth2"
	"example.com/go-web-template/restapi"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func init() {
	logLevel, _ := log.ParseLevel(AppConfig.Server.LogLevel)
	log.SetLevel(logLevel)

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{}) //for easy parsing by logstash or Splunk
}

func main() {
	gorm.DbConnect()

	// HTTP Server Set up
	// server := gin.Default() // Default Mode
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.LoggingMiddleware())
	restapi.Routes(server.Group("/restapi"))
	graphql.Routes(server.Group("/graphql"))
	oauth2.Routes(server.Group("/login"))

	server.Run("localhost:" + AppConfig.Server.PortNumber)
}
