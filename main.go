package main

import (
	"example.com/go-web-template/controller"
	"example.com/go-web-template/gorm"
	"example.com/go-web-template/graphql"
	"github.com/gin-gonic/gin"
)

func main() {
	gorm.DB_Connect()

	router := gin.Default()
	restapi.Routes(router) //Added all auth routes
	graphql.Routes(router) //Added all user routes
	router.Run()
}
