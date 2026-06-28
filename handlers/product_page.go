package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductPage(c *gin.Context) {
	var page models.ProductPage
	if err := config.DB.First(&page).Error; err != nil {
		config.DB.Create(&models.ProductPage{Description: ""})
		config.DB.First(&page)
	}
	c.JSON(http.StatusOK, gin.H{"data": page})
}

func UpdateProductPage(c *gin.Context) {
	var page models.ProductPage
	config.DB.First(&page)

	var input struct {
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	page.Description = input.Description
	config.DB.Save(&page)
	c.JSON(http.StatusOK, gin.H{"data": page})
}
