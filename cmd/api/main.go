package main

import (
	"books-api/internal/adapters/http"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	/*bookHandler := app.Init(dbcon)
	 */
	r := gin.Default()
	http.InitRoutes(r)
	if err := r.Run(":8080"); err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
}
