package controllers

import (
	"github.com/gin-gonic/gin"
	"govizpower/models"
	"net/http"
)

type ApiController struct{}

func (a *ApiController) Endpoint_kota(c *gin.Context) {
	model := models.ApiModel{}

	tahun := c.Param("tahun")
	bulan := c.Param("bulan")

	data := model.GetDataKota(tahun + bulan)

	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "0031",
			"message": "Something Wrong",
		})
	} else {
		c.JSON(http.StatusOK, data)
	}
}

func (a *ApiController) Endpoint_detail(c *gin.Context) {
	model := models.ApiModel{}
	tahun := c.Param("tahun")
	bulan := c.Param("bulan")
	idupj := c.Param("id")

	data := model.GetDetailKota(idupj, tahun+bulan)
	if data == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "0031",
			"message": "Something Wrong",
		})
	} else {
		c.JSON(http.StatusOK, data)
	}
}
