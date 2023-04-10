package main

import (
	"example.com/go-web-template/controller"
	"example.com/go-web-template/gorm"
)

func main() {
	gorm.DB_Connect()
	controller.StartRestAPI()
}
