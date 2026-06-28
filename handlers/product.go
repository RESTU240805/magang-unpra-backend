package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Preload("Images").Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProductById(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.Preload("Images").First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": product})
}

type productInput struct {
	Name          string                `json:"name"`
	Summary       string                `json:"summary"`
	Description   string                `json:"description"`
	ThumbnailPath string                `json:"thumbnail_path"`
	Category      string                `json:"category"`
	Tags          string                `json:"tags"`
	IsActive      bool                  `json:"is_active"`
	Images        []models.ProductImage `json:"Images"`
}

func CreateProduct(c *gin.Context) {
	var input productInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	product := models.Product{
		Name:          input.Name,
		Summary:       input.Summary,
		Description:   input.Description,
		ThumbnailPath: input.ThumbnailPath,
		Category:      input.Category,
		Tags:          input.Tags,
		IsActive:      input.IsActive,
	}
	config.DB.Omit("Images").Create(&product)
	for i := range input.Images {
		input.Images[i].ProductID = product.ID
		config.DB.Create(&input.Images[i])
	}
	config.DB.Preload("Images").First(&product, product.ID)
	c.JSON(http.StatusCreated, gin.H{"data": product})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")
	if err := config.DB.First(&product, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	var input map[string]interface{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	if name, ok := input["name"].(string); ok {
		product.Name = name
	}
	if summary, ok := input["summary"].(string); ok {
		product.Summary = summary
	}
	if desc, ok := input["description"].(string); ok {
		product.Description = desc
	}
	if thumb, ok := input["thumbnail_path"].(string); ok {
		product.ThumbnailPath = thumb
	}
	if cat, ok := input["category"].(string); ok {
		product.Category = cat
	}
	if tags, ok := input["tags"].(string); ok {
		product.Tags = tags
	}
	if active, ok := input["is_active"].(bool); ok {
		product.IsActive = active
	}
	config.DB.Save(&product)
	config.DB.Where("product_id = ?", product.ID).Delete(&models.ProductImage{})
	if imgs, ok := input["Images"].([]interface{}); ok {
		for _, item := range imgs {
			if imgMap, ok := item.(map[string]interface{}); ok {
				img := models.ProductImage{ProductID: product.ID}
				if url, ok := imgMap["image_url"].(string); ok {
					img.ImageURL = url
				}
				config.DB.Create(&img)
			}
		}
	}
	config.DB.Preload("Images").First(&product, product.ID)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	config.DB.Where("product_id = ?", id).Delete(&models.ProductImage{})
	config.DB.Delete(&models.Product{}, id)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}
