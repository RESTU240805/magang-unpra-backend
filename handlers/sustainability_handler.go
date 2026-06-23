package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllSustainabilities(c *gin.Context) {
	var sustainabilities []models.Sustainability

	// Mengambil data dan mengurutkannya berdasarkan sort_order terkecil
	if err := config.DB.Order("sort_order asc").Find(&sustainabilities).Error; err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sustainabilities)
}
