package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "UP",
		})
	})
}
