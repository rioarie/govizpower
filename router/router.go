package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Header", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	}
}

func Run() {
	Router := gin.Default()
	appName := "/power-api"
	version := "v0.0.1"

	Router.Use(corsMiddleware())

	apiGroup := Router.Group(appName)
	{
		apiGroup.GET("/version", func(c *gin.Context) {
			c.Writer([]byte(version))
		})
	}
}
