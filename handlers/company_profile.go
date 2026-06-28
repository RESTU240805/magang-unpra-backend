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

type profileInput struct {
	Title        string `json:"title"`
	Content      string `json:"content"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	Email        string `json:"email"`
	HeroImage    string `json:"hero_image"`
	CreedBgImage string `json:"creed_bg_image"`
}

func UpdateCompanyProfile(c *gin.Context) {
	var input profileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	var profile models.CompanyProfile
	config.DB.First(&profile)
	profile.Title = input.Title
	profile.Content = input.Content
	profile.Address = input.Address
	profile.Phone = input.Phone
	profile.Email = input.Email
	profile.HeroImage = input.HeroImage
	profile.CreedBgImage = input.CreedBgImage
	config.DB.Save(&profile)
	c.JSON(http.StatusOK, gin.H{"data": profile})
}
