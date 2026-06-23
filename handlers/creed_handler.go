package handlers

import (
	"net/http"

	"magang-unpra-backend/config"
	"magang-unpra-backend/models"

	"github.com/gin-gonic/gin"
)

func GetAllCreeds(c *gin.Context) {
	var creeds []models.Creed

	if err := config.DB.Order("sort_order asc").Find(&creeds).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, creeds)
}

func GetCreedById(c *gin.Context) {
	id := c.Param("id")
	var creed models.Creed

	if err := config.DB.First(&creed, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Creed tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, creed)
}

func CreateCreed(c *gin.Context) {
	var input struct {
		TitleJP     string `json:"title_jp"`
		TitleEN     string `json:"title_en" binding:"required"`
		Roma        string `json:"roma"`
		Tagline     string `json:"tagline"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
		IsActive    bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	creed := models.Creed{
		TitleJP:     input.TitleJP,
		TitleEN:     input.TitleEN,
		Roma:        input.Roma,
		Tagline:     input.Tagline,
		Description: input.Description,
		SortOrder:   input.SortOrder,
		IsActive:    input.IsActive,
	}

	if err := config.DB.Create(&creed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, creed)
}

func UpdateCreed(c *gin.Context) {
	id := c.Param("id")
	var creed models.Creed

	if err := config.DB.First(&creed, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Creed tidak ditemukan",
		})
		return
	}

	var input struct {
		TitleJP     string `json:"title_jp"`
		TitleEN     string `json:"title_en"`
		Roma        string `json:"roma"`
		Tagline     string `json:"tagline"`
		Description string `json:"description"`
		SortOrder   int    `json:"sort_order"`
		IsActive    bool   `json:"is_active"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	creed.TitleJP = input.TitleJP
	creed.TitleEN = input.TitleEN
	creed.Roma = input.Roma
	creed.Tagline = input.Tagline
	creed.Description = input.Description
	creed.SortOrder = input.SortOrder
	creed.IsActive = input.IsActive

	if err := config.DB.Save(&creed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, creed)
}

func DeleteCreed(c *gin.Context) {
	id := c.Param("id")
	var creed models.Creed

	if err := config.DB.First(&creed, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Creed tidak ditemukan",
		})
		return
	}

	if err := config.DB.Delete(&creed).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Creed berhasil dihapus",
	})
}
