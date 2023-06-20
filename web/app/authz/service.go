package authz

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(rg *gin.RouterGroup) {
	rg.GET("/authorized", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	rg.GET("/unauthorized", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})
}
