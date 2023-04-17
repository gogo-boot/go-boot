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

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	gorm.DbConnect()
	//server := gin.Default()
	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middlewares.LoggingMiddleware())

	restapi.Routes(server.Group("/restapi"))
	graphql.Routes(server.Group("/graphql"))
	oauth2.Routes(server.Group("/login"))

	server.Run("localhost:" + AppConfig.Server.PortNumber)
}
