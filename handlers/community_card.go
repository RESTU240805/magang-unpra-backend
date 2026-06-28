package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllCommunityCards(c *gin.Context) {
	var cards []models.CommunityCard
	config.DB.Preload("Images").Where("is_active = ?", true).Order("sort_order ASC, id ASC").Find(&cards)
	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func GetAllCommunityCardsAdmin(c *gin.Context) {
	var cards []models.CommunityCard
	config.DB.Preload("Images").Order("sort_order ASC, id ASC").Find(&cards)
	c.JSON(http.StatusOK, gin.H{"data": cards})
}

func CreateCommunityCard(c *gin.Context) {
	var input struct {
		Title       string                      `json:"title"`
		Description string                      `json:"description"`
		IconPath    string                      `json:"icon_path"`
		Link        string                      `json:"link"`
		SortOrder   int                         `json:"sort_order"`
		IsActive    bool                        `json:"is_active"`
		Images      []models.CommunityCardImage `json:"images"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	if input.Link == "" {
		input.Link = "/sustainability/csr/community"
	}

	card := models.CommunityCard{
		Title:       input.Title,
		Description: input.Description,
		IconPath:    input.IconPath,
		Link:        input.Link,
		SortOrder:   input.SortOrder,
		IsActive:    input.IsActive,
	}
	config.DB.Create(&card)

	for i := range input.Images {
		input.Images[i].CardID = card.ID
		config.DB.Create(&input.Images[i])
	}

	config.DB.Preload("Images").First(&card, card.ID)
	c.JSON(http.StatusOK, gin.H{"data": card})
}

func UpdateCommunityCard(c *gin.Context) {
	id := c.Param("id")
	var card models.CommunityCard
	if err := config.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}

	var input struct {
		Title       string                      `json:"title"`
		Description string                      `json:"description"`
		IconPath    string                      `json:"icon_path"`
		Link        string                      `json:"link"`
		SortOrder   int                         `json:"sort_order"`
		IsActive    bool                        `json:"is_active"`
		Images      []models.CommunityCardImage `json:"images"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	card.Title = input.Title
	card.Description = input.Description
	card.IconPath = input.IconPath
	card.Link = input.Link
	card.SortOrder = input.SortOrder
	card.IsActive = input.IsActive
	config.DB.Save(&card)

	config.DB.Where("card_id = ?", card.ID).Delete(&models.CommunityCardImage{})
	for i := range input.Images {
		input.Images[i].CardID = card.ID
		config.DB.Create(&input.Images[i])
	}

	config.DB.Preload("Images").First(&card, card.ID)
	c.JSON(http.StatusOK, gin.H{"data": card})
}

func DeleteCommunityCard(c *gin.Context) {
	id := c.Param("id")
	var card models.CommunityCard
	if err := config.DB.First(&card, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Card not found"})
		return
	}
	config.DB.Where("card_id = ?", card.ID).Delete(&models.CommunityCardImage{})
	config.DB.Delete(&card)
	c.JSON(http.StatusOK, gin.H{"message": "Card deleted"})
}
