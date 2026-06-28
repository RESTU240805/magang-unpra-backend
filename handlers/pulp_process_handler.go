package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ─── Pulp Process Sections ─────────────────────────────────

func GetAllPulpProcessSections(c *gin.Context) {
	var items []models.PulpProcessSection
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllPulpProcessSectionsAdmin(c *gin.Context) {
	var items []models.PulpProcessSection
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type pulpProcessSectionInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	SortOrder   int    `json:"sort_order"`
	IsActive    *bool  `json:"is_active"`
}

func CreatePulpProcessSection(c *gin.Context) {
	var input pulpProcessSectionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.PulpProcessSection{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		SortOrder:   input.SortOrder,
		IsActive:    true,
	}
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdatePulpProcessSection(c *gin.Context) {
	var item models.PulpProcessSection
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Section not found"})
		return
	}
	var input pulpProcessSectionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.Title = input.Title
	item.Description = input.Description
	item.ImageURL = input.ImageURL
	item.SortOrder = input.SortOrder
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeletePulpProcessSection(c *gin.Context) {
	config.DB.Delete(&models.PulpProcessSection{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Section deleted"})
}

// ─── Pulp Process Recovery Tabs ────────────────────────────

func GetAllPulpProcessRecoveries(c *gin.Context) {
	var items []models.PulpProcessRecovery
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllPulpProcessRecoveriesAdmin(c *gin.Context) {
	var items []models.PulpProcessRecovery
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type pulpProcessRecoveryInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
	Specs       string `json:"specs"`
	SortOrder   int    `json:"sort_order"`
	IsActive    *bool  `json:"is_active"`
}

func CreatePulpProcessRecovery(c *gin.Context) {
	var input pulpProcessRecoveryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.PulpProcessRecovery{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		Specs:       input.Specs,
		SortOrder:   input.SortOrder,
		IsActive:    true,
	}
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdatePulpProcessRecovery(c *gin.Context) {
	var item models.PulpProcessRecovery
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recovery not found"})
		return
	}
	var input pulpProcessRecoveryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.Title = input.Title
	item.Description = input.Description
	item.ImageURL = input.ImageURL
	item.Specs = input.Specs
	item.SortOrder = input.SortOrder
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeletePulpProcessRecovery(c *gin.Context) {
	config.DB.Delete(&models.PulpProcessRecovery{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Recovery deleted"})
}
