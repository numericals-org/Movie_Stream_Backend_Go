package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	message := map[string]any{"message": "hello world"}

	r.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, message)
	})

	r.Run("localhost:8080")
}
