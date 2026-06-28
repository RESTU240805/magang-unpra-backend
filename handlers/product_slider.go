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

type slideInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	SortOrder   int    `json:"sort_order"`
}

func CreateSlide(c *gin.Context) {
	var input slideInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	slide := models.ProductSlider{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		SortOrder:   input.SortOrder,
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
	var input slideInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
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
