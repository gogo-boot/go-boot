package actuator

import (
	"github.com/gin-gonic/gin"
	myConfig "gogo-boot/go-boot/platform/config"
	"net/http"
)

func config(c *gin.Context) {
	c.JSON(http.StatusOK, myConfig.AppConfig)
}
