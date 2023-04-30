package actuator

import (
	"example.com/go-boot/platform/initializer"
	"github.com/gin-gonic/gin"
)

func init() {
	Routes(initializer.Router.Group("/actuator"))
}

func Routes(rg *gin.RouterGroup) {
	rg.GET("/health", health)
	rg.GET("/mem", mem)
	// Todo mask security relevant properties
	rg.GET("/config", config)
}
