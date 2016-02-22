package controllers

import (
	"github.com/gin-gonic/gin"
	"govizpower/models"
	"net/http"
)

type ApiController struct{}

// Endpoint_kota is handler for request from /all/:tahun/:bulan
func (a *ApiController) Endpoint_kota(c *gin.Context) {
	model := models.ApiModel{}

	tahun := c.Param("tahun") // year
	bulan := c.Param("bulan") // month

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

// Endpoint_detail is handler for request from /all/:tahun/:bulan/:id
func (a *ApiController) Endpoint_detail(c *gin.Context) {
	model := models.ApiModel{}
	tahun := c.Param("tahun") //year
	bulan := c.Param("bulan") // month
	idupj := c.Param("id") // idcity

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
