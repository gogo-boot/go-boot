package actuator

import "github.com/gin-gonic/gin"

func Routes(rg *gin.RouterGroup) {
	rg.GET("/health", health)
	rg.GET("/mem", mem)
	rg.GET("/config", config)
}
