package main

import (
	"example.com/go-web-template/gorm"
	"example.com/go-web-template/graphql"
	"example.com/go-web-template/oauth2"
	"example.com/go-web-template/restapi"
	"github.com/gin-gonic/gin"
)

func main() {
	gorm.DbConnect()

	router := gin.Default()
	restapi.Routes(router)
	graphql.Routes(router)
	oauth2.Routes(router)
	router.Run()
}
