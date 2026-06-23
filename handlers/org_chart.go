package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrgChart(c *gin.Context) {
	var chart models.OrgChart
	result := config.DB.First(&chart)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": chart})
}

func UpdateOrgChart(c *gin.Context) {
	var input struct {
		ImagePath string `json:"image_path"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var chart models.OrgChart
	result := config.DB.First(&chart)
	if result.Error != nil {
		chart = models.OrgChart{ImagePath: input.ImagePath}
		config.DB.Create(&chart)
	} else {
		chart.ImagePath = input.ImagePath
		config.DB.Save(&chart)
	}

	c.JSON(http.StatusOK, gin.H{"data": chart})
}
