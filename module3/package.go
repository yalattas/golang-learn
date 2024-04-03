package module3

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ModuleTest() {
	// Initialize the Gin router
	router := gin.Default()

	// Define a route handler
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Gin!")
	})

	// Run the server
	router.Run(":8080")
}
