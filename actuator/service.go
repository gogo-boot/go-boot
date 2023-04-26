package actuator

import "github.com/gin-gonic/gin"

func Routes(rg *gin.RouterGroup) {
	rg.GET("/health", health)
	rg.GET("/mem", mem)
	// Todo mask security relevant properties
	rg.GET("/config", config)
}
