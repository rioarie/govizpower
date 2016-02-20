package controllers

import (
	"github.com/gin-gonic/gin"
	_ "govizpower/datamappers"
	"govizpower/models"
	"net/http"
)

type ApiController struct{}

func (a *ApiController) Endpoint_kota(c *gin.Context) {
	model := models.ApiModel{}

	data := model.GetDataKota()

	// if data == (datamappers.KotaMapper{}) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  "0031",
	// 		"message": "Someting Wrong",
	// 	})
	// } else {
	c.JSON(http.StatusOK, data)
	// }
}
