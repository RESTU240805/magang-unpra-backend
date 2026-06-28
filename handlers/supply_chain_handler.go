package handlers

import (
	"encoding/json"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func toJSONString(v interface{}) string {
	switch val := v.(type) {
	case string:
		return val
	case []byte:
		return string(val)
	default:
		b, _ := json.Marshal(val)
		return string(b)
	}
}

// ─── Supply Chain Strategy Items ──────────────────────────

func GetAllSupplyChainStrategies(c *gin.Context) {
	var items []models.SupplyChainStrategy
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllSupplyChainStrategiesAdmin(c *gin.Context) {
	var items []models.SupplyChainStrategy
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type supplyChainStrategyInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	IsActive    *bool  `json:"is_active"`
}

func CreateSupplyChainStrategy(c *gin.Context) {
	var input supplyChainStrategyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.SupplyChainStrategy{
		Title:       input.Title,
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

func UpdateSupplyChainStrategy(c *gin.Context) {
	var item models.SupplyChainStrategy
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Strategy not found"})
		return
	}
	var input supplyChainStrategyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.Title = input.Title
	item.Description = input.Description
	item.SortOrder = input.SortOrder
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteSupplyChainStrategy(c *gin.Context) {
	config.DB.Delete(&models.SupplyChainStrategy{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Strategy deleted"})
}

// ─── Supply Chain Sustainability Items ────────────────────

func GetAllSupplyChainSustainabilityItems(c *gin.Context) {
	var items []models.SupplyChainSustainabilityItem
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllSupplyChainSustainabilityItemsAdmin(c *gin.Context) {
	var items []models.SupplyChainSustainabilityItem
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type supplyChainSustainabilityInput struct {
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
	IsActive    *bool  `json:"is_active"`
}

func CreateSupplyChainSustainabilityItem(c *gin.Context) {
	var input supplyChainSustainabilityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.SupplyChainSustainabilityItem{
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

func UpdateSupplyChainSustainabilityItem(c *gin.Context) {
	var item models.SupplyChainSustainabilityItem
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		return
	}
	var input supplyChainSustainabilityInput
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

func DeleteSupplyChainSustainabilityItem(c *gin.Context) {
	config.DB.Delete(&models.SupplyChainSustainabilityItem{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
}

// ─── Supply Chain Policies ────────────────────────────────

func GetAllSupplyChainPolicies(c *gin.Context) {
	var items []models.SupplyChainPolicy
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllSupplyChainPoliciesAdmin(c *gin.Context) {
	var items []models.SupplyChainPolicy
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type supplyChainPolicyInput struct {
	Title      string          `json:"title"`
	Points     json.RawMessage `json:"points"`
	Procedures json.RawMessage `json:"procedures"`
	SortOrder  int             `json:"sort_order"`
	IsActive   *bool           `json:"is_active"`
}

func CreateSupplyChainPolicy(c *gin.Context) {
	var input supplyChainPolicyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.SupplyChainPolicy{
		Title:      input.Title,
		Points:     toJSONString(input.Points),
		Procedures: toJSONString(input.Procedures),
		SortOrder:  input.SortOrder,
		IsActive:   true,
	}
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Create(&item)
	c.JSON(http.StatusCreated, gin.H{"data": item})
}

func UpdateSupplyChainPolicy(c *gin.Context) {
	var item models.SupplyChainPolicy
	id := c.Param("id")
	if err := config.DB.First(&item, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
		return
	}
	var input supplyChainPolicyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item.Title = input.Title
	item.Points = toJSONString(input.Points)
	item.Procedures = toJSONString(input.Procedures)
	item.SortOrder = input.SortOrder
	if input.IsActive != nil {
		item.IsActive = *input.IsActive
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteSupplyChainPolicy(c *gin.Context) {
	config.DB.Delete(&models.SupplyChainPolicy{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Policy deleted"})
}
