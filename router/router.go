package router

import (
	"github.com/gin-gonic/gin"
	"govizpower/controllers"
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
	version := "v0.0.2"

	Router.Use(corsMiddleware())

	apiCtrl := controllers.ApiController{}

	apiGroup := Router.Group(appName)
	{
		apiGroup.GET("/all/:tahun/:bulan", apiCtrl.Endpoint_kota)
		apiGroup.GET("/detail/:tahun/:bulan/:id", apiCtrl.Endpoint_detail)
		apiGroup.GET("/version", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"App Name": "Go Viz Power",
				"Version":  version,
			})
		})
	}

	Router.Run(":4040")
}
