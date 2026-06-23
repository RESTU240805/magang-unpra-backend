package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCompanyProfile(c *gin.Context) {
	var profile models.CompanyProfile
	config.DB.First(&profile)
	c.JSON(http.StatusOK, gin.H{"data": profile})
}

func UpdateCompanyProfile(c *gin.Context) {
	var profile models.CompanyProfile
	config.DB.First(&profile)
	c.ShouldBindJSON(&profile)
	config.DB.Save(&profile)
	c.JSON(http.StatusOK, gin.H{"data": profile})
}
