package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ─── Wood Types ──────────────────────────────────────────────

func GetAllWoodTypes(c *gin.Context) {
	var items []models.ForestWoodType
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type woodTypeInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Badge       string `json:"badge"`
	SortOrder   int    `json:"sort_order"`
}

func CreateWoodType(c *gin.Context) {
	var input woodTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.ForestWoodType{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Badge:       input.Badge,
		SortOrder:   input.SortOrder,
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdateWoodType(c *gin.Context) {
	var item models.ForestWoodType
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Wood type not found"})
		return
	}
	var input woodTypeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.Title = input.Title
	item.Description = input.Description
	item.ImageURL = input.ImageURL
	item.Badge = input.Badge
	item.SortOrder = input.SortOrder
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteWoodType(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.ForestWoodType{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Wood type deleted"})
}

// ─── Approach ────────────────────────────────────────────────

func GetForestApproach(c *gin.Context) {
	var approach models.ForestApproach
	if err := config.DB.First(&approach).Error; err != nil {
		config.DB.Create(&models.ForestApproach{Description: ""})
		config.DB.First(&approach)
	}
	c.JSON(http.StatusOK, gin.H{"data": approach})
}

func UpdateForestApproach(c *gin.Context) {
	var approach models.ForestApproach
	config.DB.First(&approach)

	var input struct {
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	approach.Description = input.Description
	config.DB.Save(&approach)
	c.JSON(http.StatusOK, gin.H{"data": approach})
}

// ─── Sliders ─────────────────────────────────────────────────

func GetAllForestSliders(c *gin.Context) {
	var slides []models.ForestSlider
	config.DB.Order("sort_order asc").Find(&slides)
	c.JSON(http.StatusOK, gin.H{"data": slides})
}

type forestSliderInput struct {
	ImageURL  string `json:"image_url"`
	SortOrder int    `json:"sort_order"`
}

func CreateForestSlider(c *gin.Context) {
	var input forestSliderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	slide := models.ForestSlider{
		ImageURL:  input.ImageURL,
		SortOrder: input.SortOrder,
	}
	config.DB.Create(&slide)
	c.JSON(http.StatusCreated, gin.H{"data": slide})
}

func UpdateForestSlider(c *gin.Context) {
	var slide models.ForestSlider
	id := c.Param("id")
	if err := config.DB.First(&slide, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slider not found"})
		return
	}
	var input forestSliderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	slide.ImageURL = input.ImageURL
	slide.SortOrder = input.SortOrder
	config.DB.Save(&slide)
	c.JSON(http.StatusOK, gin.H{"data": slide})
}

func DeleteForestSlider(c *gin.Context) {
	id := c.Param("id")
	config.DB.Delete(&models.ForestSlider{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Slider deleted"})
}
