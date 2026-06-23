package handlers

import (
	"fmt"
	"io"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func UploadProductImage(c *gin.Context) {
	file, header, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}
	defer file.Close()

	ext := strings.ToLower(filepath.Ext(header.Filename))
	allowed := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".webp": true}
	if !allowed[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung"})
		return
	}

	uploadDir := "./uploads/products"
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat folder"})
		return
	}

	filename := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), uuid.New().String()[:8], ext)
	savePath := filepath.Join(uploadDir, filename)

	dst, err := os.Create(savePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menulis file"})
		return
	}

	imageURL := "/uploads/products/" + filename
	c.JSON(http.StatusOK, gin.H{"url": imageURL})
}

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

func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Omit("Images").Create(&product)
	for i := range product.Images {
		product.Images[i].ProductID = product.ID
		config.DB.Create(&product.Images[i])
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
