package actuator

import (
	myConfig "example.com/go-boot/platform/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func config(c *gin.Context) {
	c.JSON(http.StatusOK, myConfig.AppConfig)
}
