package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllSlides(c *gin.Context) {
	var slides []models.ProductSlider
	config.DB.Order("sort_order asc").Find(&slides)
	c.JSON(http.StatusOK, gin.H{"data": slides})
}

func CreateSlide(c *gin.Context) {
	var slide models.ProductSlider
	if err := c.ShouldBindJSON(&slide); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&slide)
	c.JSON(http.StatusCreated, gin.H{"data": slide})
}

func UpdateSlide(c *gin.Context) {
	var slide models.ProductSlider
	id := c.Param("id")
	if err := config.DB.First(&slide, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slide not found"})
		return
	}
	var input models.ProductSlider
	c.ShouldBindJSON(&input)
	slide.Title = input.Title
	slide.Description = input.Description
	slide.ImageURL = input.ImageURL
	slide.SortOrder = input.SortOrder
	config.DB.Save(&slide)
	c.JSON(http.StatusOK, gin.H{"data": slide})
}

func DeleteSlide(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.ProductSlider{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Slide deleted"})
}
