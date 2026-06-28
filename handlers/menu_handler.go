package handlers

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuResponse struct {
	ID        uint           `json:"id"`
	Name      string         `json:"name"`
	URL       string         `json:"url"`
	IsActive  bool           `json:"is_active"`
	ParentID  *uint          `json:"parent_id"`
	SortOrder int            `json:"sort_order"`
	Children  []MenuResponse `json:"children,omitempty"`
}

type MenuInput struct {
	Name      string `json:"name"`
	URL       string `json:"url"`
	IsActive  bool   `json:"is_active"`
	ParentID  *uint  `json:"parent_id"`
	SortOrder int    `json:"sort_order"`
}

func buildMenuTree(menus []models.Menu, parentID *uint) []MenuResponse {
	var result []MenuResponse
	for _, m := range menus {
		if (parentID == nil && m.ParentID == nil) || (parentID != nil && m.ParentID != nil && *m.ParentID == *parentID) {
			children := buildMenuTree(menus, &m.ID)
			item := MenuResponse{
				ID:        m.ID,
				Name:      m.Name,
				URL:       m.URL,
				IsActive:  m.IsActive,
				ParentID:  m.ParentID,
				SortOrder: m.SortOrder,
			}
			if len(children) > 0 {
				item.Children = children
			}
			result = append(result, item)
		}
	}
	return result
}

func GetActiveMenus(c *gin.Context) {
	var menus []models.Menu
	config.DB.Where("is_active = ?", true).Order("sort_order ASC, id ASC").Find(&menus)
	tree := buildMenuTree(menus, nil)
	c.JSON(http.StatusOK, gin.H{"data": tree})
}

func GetAllMenus(c *gin.Context) {
	var menus []models.Menu
	config.DB.Order("sort_order ASC, id ASC").Find(&menus)
	tree := buildMenuTree(menus, nil)
	c.JSON(http.StatusOK, gin.H{"data": tree})
}

func CreateMenu(c *gin.Context) {
	var input MenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	menu := models.Menu{
		Name:      input.Name,
		URL:       input.URL,
		IsActive:  input.IsActive,
		ParentID:  input.ParentID,
		SortOrder: input.SortOrder,
	}

	if err := config.DB.Create(&menu).Error; err != nil {
		log.Printf("Failed to create menu: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan menu"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": menu})
}

func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	if err := config.DB.First(&menu, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Menu not found"})
		return
	}

	var input MenuInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	menu.Name = input.Name
	menu.URL = input.URL
	menu.IsActive = input.IsActive
	menu.ParentID = input.ParentID
	menu.SortOrder = input.SortOrder

	if err := config.DB.Save(&menu).Error; err != nil {
		log.Printf("Failed to update menu: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": menu})
}

func DeleteMenu(c *gin.Context) {
	if err := config.DB.Delete(&models.Menu{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Menu deleted"})
}
