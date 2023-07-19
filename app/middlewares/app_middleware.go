package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starter-kit-go/config"
)

func AppMiddleware(c *gin.Context) {
	appKey := c.GetHeader("app-key")
	if config.App().Key != appKey {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "you can't do this!",
		})
		return
	}
}
