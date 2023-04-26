package actuator

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

func mem(c *gin.Context) {
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	c.JSON(http.StatusOK, m)
}
