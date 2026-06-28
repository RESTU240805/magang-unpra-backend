package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAboutSection(c *gin.Context) {
	var about models.AboutSection
	if err := config.DB.First(&about).Error; err != nil {
		config.DB.Create(&models.AboutSection{
			Title:       "ABOUT TELPP",
			Description: "PT Tanjungenim Lestari Pulp and Paper (TELPP) is world class manufacturer of high product quality and environmental friendly market pulp mill.",
			ImagePath:   "/images/gedung.jpeg",
			BadgeNumber: "20+",
			BadgeLabel:  "Years of Excellence",
		})
		config.DB.First(&about)
	}
	c.JSON(http.StatusOK, gin.H{"data": about})
}

func UpdateAboutSection(c *gin.Context) {
	var about models.AboutSection
	config.DB.First(&about)

	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		ImagePath   string `json:"image_path"`
		BadgeNumber string `json:"badge_number"`
		BadgeLabel  string `json:"badge_label"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	about.Title = input.Title
	about.Description = input.Description
	about.ImagePath = input.ImagePath
	about.BadgeNumber = input.BadgeNumber
	about.BadgeLabel = input.BadgeLabel
	config.DB.Save(&about)
	c.JSON(http.StatusOK, gin.H{"data": about})
}
