package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"
	"strings"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
)

func makeSlug(title string) string {
	var result strings.Builder
	for _, r := range strings.ToLower(title) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			result.WriteRune(r)
		} else if unicode.IsSpace(r) || r == '-' || r == '_' {
			result.WriteRune('-')
		}
	}
	return result.String()
}

func makeUniqueSlug(title string, excludeID uint) string {
	slug := makeSlug(title)
	if slug == "" {
		slug = "untitled"
	}
	var existing models.News
	for config.DB.Where("slug = ?", slug).Not("id = ?", excludeID).First(&existing).Error == nil {
		slug = slug + "-1"
	}
	return slug
}

// GET /api/news — hanya yang published, untuk halaman publik
func GetAllNews(c *gin.Context) {
	var news []models.News
	config.DB.Preload("Images").Where("is_published = ?", true).Find(&news)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

// GET /api/admin/news — semua berita (admin)
func GetAllNewsAdmin(c *gin.Context) {
	var news []models.News
	config.DB.Preload("Images").Find(&news)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

// GET /api/news/:id
func GetNewsById(c *gin.Context) {
	var news models.News
	id := c.Param("id")
	if err := config.DB.Preload("Images").First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": news})
}

// POST /api/admin/news
func CreateNews(c *gin.Context) {
	var input struct {
		Title         string             `json:"title"          binding:"required"`
		Slug          string             `json:"slug"`
		Summary       string             `json:"summary"`
		Content       string             `json:"content"        binding:"required"`
		Category      string             `json:"category"`
		ThumbnailPath string             `json:"thumbnail_path"`
		IsPublished   bool               `json:"is_published"`
		PublishedAt   *string            `json:"published_at"`
		Images        []models.NewsImage `json:"Images"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	news := models.News{
		Title:         input.Title,
		Slug:          makeUniqueSlug(input.Title, 0),
		Summary:       input.Summary,
		Content:       input.Content,
		Category:      input.Category,
		ThumbnailPath: input.ThumbnailPath,
		IsPublished:   input.IsPublished,
	}

	if input.PublishedAt != nil && *input.PublishedAt != "" {
		parsed, err := time.Parse("2006-01-02", *input.PublishedAt)
		if err != nil {
			parsed, err = time.Parse(time.RFC3339, *input.PublishedAt)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal tidak valid. Gunakan YYYY-MM-DD"})
				return
			}
		}
		news.PublishedAt = &parsed
	} else {
		now := time.Now()
		news.PublishedAt = &now
	}

	if err := config.DB.Omit("Images").Create(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan berita"})
		return
	}

	for i := range input.Images {
		input.Images[i].NewsID = news.ID
		config.DB.Create(&input.Images[i])
	}

	config.DB.Preload("Images").First(&news, news.ID)
	c.JSON(http.StatusCreated, gin.H{"data": news})
}

// PUT /api/admin/news/:id
func UpdateNews(c *gin.Context) {
	id := c.Param("id")

	var news models.News
	if err := config.DB.First(&news, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "News not found"})
		return
	}

	var input struct {
		Title         string             `json:"title"`
		Slug          string             `json:"slug"`
		Summary       string             `json:"summary"`
		Content       string             `json:"content"`
		Category      string             `json:"category"`
		ThumbnailPath string             `json:"thumbnail_path"`
		IsPublished   bool               `json:"is_published"`
		PublishedAt   *string            `json:"published_at"`
		Images        []models.NewsImage `json:"Images"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	// Update semua field termasuk thumbnail_path & is_published
	news.Title = input.Title
	news.Summary = input.Summary
	news.Content = input.Content
	news.Category = input.Category
	news.ThumbnailPath = input.ThumbnailPath
	news.IsPublished = input.IsPublished

	if input.Slug != "" {
		news.Slug = makeUniqueSlug(input.Slug, news.ID)
	} else {
		news.Slug = makeUniqueSlug(input.Title, news.ID)
	}

	// Jika admin mengisi tanggal baru, update; jika tidak diisi, biarkan tanggal lama
	if input.PublishedAt != nil && *input.PublishedAt != "" {
		parsed, err := time.Parse("2006-01-02", *input.PublishedAt)
		if err != nil {
			parsed, err = time.Parse(time.RFC3339, *input.PublishedAt)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Format tanggal tidak valid. Gunakan YYYY-MM-DD"})
				return
			}
		}
		news.PublishedAt = &parsed
	}

	if err := config.DB.Save(&news).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan berita"})
		return
	}

	config.DB.Where("news_id = ?", news.ID).Delete(&models.NewsImage{})
	for i := range input.Images {
		input.Images[i].NewsID = news.ID
		config.DB.Create(&input.Images[i])
	}

	config.DB.Preload("Images").First(&news, news.ID)
	c.JSON(http.StatusOK, gin.H{"data": news})
}

// DELETE /api/admin/news/:id
func DeleteNews(c *gin.Context) {
	id := c.Param("id")
	config.DB.Where("news_id = ?", id).Delete(&models.NewsImage{})
	if err := config.DB.Delete(&models.News{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus berita"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "News deleted"})
}
