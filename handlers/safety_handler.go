package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ─── Safety Policies ──────────────────────────────────────

func GetAllSafetyPolicies(c *gin.Context) {
	var items []models.SafetyPolicy
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Limit(1).Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllSafetyPoliciesAdmin(c *gin.Context) {
	var items []models.SafetyPolicy
	config.DB.Order("sort_order asc").Limit(1).Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type safetyPolicyInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Icon        *string `json:"icon"`
	SortOrder   *int    `json:"sort_order"`
	IsActive    *bool   `json:"is_active"`
}

func CreateSafetyPolicy(c *gin.Context) {
	var input safetyPolicyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.SafetyPolicy{
		IsActive: true,
	}
	if input.Title != nil {
		item.Title = *input.Title
	}
	if input.Description != nil {
		item.Description = *input.Description
	}
	if input.Icon != nil {
		item.Icon = *input.Icon
	}
	if input.SortOrder != nil {
		item.SortOrder = *input.SortOrder
	}
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdateSafetyPolicy(c *gin.Context) {
	id := c.Param("id")
	var input safetyPolicyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	updates := map[string]interface{}{}
	if input.Title != nil {
		updates["title"] = *input.Title
	}
	if input.Description != nil {
		updates["description"] = *input.Description
	}
	if input.Icon != nil {
		updates["icon"] = *input.Icon
	}
	if input.SortOrder != nil {
		updates["sort_order"] = *input.SortOrder
	}
	if input.IsActive != nil {
		updates["is_active"] = *input.IsActive
	}
	if len(updates) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tidak ada data yang diupdate"})
		return
	}
	result := config.DB.Model(&models.SafetyPolicy{}).Where("id = ?", id).Updates(updates)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
		return
	}
	var item models.SafetyPolicy
	config.DB.First(&item, id)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteSafetyPolicy(c *gin.Context) {
	config.DB.Delete(&models.SafetyPolicy{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Policy deleted"})
}

// ─── Safety K3 Targets ────────────────────────────────────

func GetAllSafetyK3Targets(c *gin.Context) {
	var items []models.SafetyK3Target
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllSafetyK3TargetsAdmin(c *gin.Context) {
	var items []models.SafetyK3Target
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type safetyK3TargetInput struct {
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	IsActive    *bool  `json:"is_active"`
}

func CreateSafetyK3Target(c *gin.Context) {
	var input safetyK3TargetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.SafetyK3Target{
		Description: input.Description,
		SortOrder:   input.SortOrder,
		IsActive:    true,
	}
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdateSafetyK3Target(c *gin.Context) {
	var item models.SafetyK3Target
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Target not found"})
		return
	}
	var input safetyK3TargetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.Description = input.Description
	item.SortOrder = input.SortOrder
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteSafetyK3Target(c *gin.Context) {
	config.DB.Delete(&models.SafetyK3Target{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Target deleted"})
}

// ─── Safety K3 Programs ───────────────────────────────────

func GetAllSafetyK3Programs(c *gin.Context) {
	var items []models.SafetyK3Program
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllSafetyK3ProgramsAdmin(c *gin.Context) {
	var items []models.SafetyK3Program
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type safetyK3ProgramInput struct {
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	IsActive    *bool  `json:"is_active"`
}

func CreateSafetyK3Program(c *gin.Context) {
	var input safetyK3ProgramInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.SafetyK3Program{
		Description: input.Description,
		SortOrder:   input.SortOrder,
		IsActive:    true,
	}
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdateSafetyK3Program(c *gin.Context) {
	var item models.SafetyK3Program
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Program not found"})
		return
	}
	var input safetyK3ProgramInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.Description = input.Description
	item.SortOrder = input.SortOrder
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteSafetyK3Program(c *gin.Context) {
	config.DB.Delete(&models.SafetyK3Program{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Program deleted"})
}

// ─── Safety Sliders ───────────────────────────────────────

func GetAllSafetySliders(c *gin.Context) {
	var items []models.SafetySlider
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllSafetySlidersAdmin(c *gin.Context) {
	var items []models.SafetySlider
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type safetySliderInput struct {
	ImageURL  string `json:"image_url"`
	SortOrder int    `json:"sort_order"`
	IsActive  *bool  `json:"is_active"`
}

func CreateSafetySlider(c *gin.Context) {
	var input safetySliderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.SafetySlider{
		ImageURL:  input.ImageURL,
		SortOrder: input.SortOrder,
		IsActive:  true,
	}
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdateSafetySlider(c *gin.Context) {
	var item models.SafetySlider
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slider not found"})
		return
	}
	var input safetySliderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.ImageURL = input.ImageURL
	item.SortOrder = input.SortOrder
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteSafetySlider(c *gin.Context) {
	config.DB.Delete(&models.SafetySlider{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Slider deleted"})
}
