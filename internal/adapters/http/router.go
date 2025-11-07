package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes(r *gin.Engine, bookHandler *BookHandler) {
	r.GET("/books/:id", bookHandler.handleGetBook)
	r.GET("/books", bookHandler.handleGetAllBooks)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "UP",
		})
	})
}
