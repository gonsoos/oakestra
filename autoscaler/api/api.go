package api

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	port = os.Getenv("API_PORT")
)

// StartApiServer starts the Gin HTTP server.
// It returns an error if the HTTP server fails.
func StartApiServer() error {
	router := gin.Default()
	router.GET("/status", getStatus)

	return router.Run(fmt.Sprintf(":%v", port))
}

func getStatus(c *gin.Context) {
	c.Status(http.StatusOK)
}
