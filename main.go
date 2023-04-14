package main

import (
	. "example.com/go-web-template/config"
	"example.com/go-web-template/gorm"
	"example.com/go-web-template/graphql"
	"example.com/go-web-template/oauth2"
	"example.com/go-web-template/restapi"
	"github.com/gin-gonic/gin"
)

func main() {
	gorm.DbConnect()
	r := gin.Default()
	restapi.Routes(r.Group("/restapi"))
	graphql.Routes(r.Group("/graphql"))
	oauth2.Routes(r.Group("/login"))
	r.Run(":" + AppConfig.Server.PortNumber)
}
