package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"magang-unpra-backend/config"
	"magang-unpra-backend/models"

	"github.com/gin-gonic/gin"
)

// GET /api/company-documents (publik, untuk frontend Vue)
func GetAllCompanyDocuments(c *gin.Context) {
	var docs []models.CompanyDocument
	if err := config.DB.Order("created_at desc").Find(&docs).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data"})
		return
	}
	c.JSON(http.StatusOK, docs)
}

// POST /api/admin/company-documents
func CreateCompanyDocument(c *gin.Context) {
	title := c.PostForm("title")
	category := c.PostForm("category")
	docDate := c.PostForm("doc_date")
	description := c.PostForm("description")

	var fileURL, fileType, fileSize string

	file, err := c.FormFile("file")
	if err == nil {
		if file.Size > 10<<20 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File terlalu besar. Maksimal 10 MB"})
			return
		}

		ext := filepath.Ext(file.Filename)
		allowedDocExts := map[string]bool{".pdf": true, ".doc": true, ".docx": true, ".xls": true, ".xlsx": true, ".ppt": true, ".pptx": true}
		if !allowedDocExts[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung. Gunakan PDF, DOC, XLS, atau PPT"})
			return
		}
		fileType = ext[1:]

		filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
			return
		}

		fileURL = "/uploads/" + filename
		sizeKB := float64(file.Size) / 1024
		if sizeKB > 1024 {
			fileSize = strconv.FormatFloat(sizeKB/1024, 'f', 1, 64) + " MB"
		} else {
			fileSize = strconv.FormatFloat(sizeKB, 'f', 0, 64) + " KB"
		}
	}

	doc := models.CompanyDocument{
		Title:       title,
		Category:    category,
		DocDate:     docDate,
		FileType:    fileType,
		FileSize:    fileSize,
		Description: description,
		FileURL:     fileURL,
	}

	if err := config.DB.Create(&doc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan data"})
		return
	}

	c.JSON(http.StatusCreated, doc)
}

// PUT /api/admin/company-documents/:id
func UpdateCompanyDocument(c *gin.Context) {
	id := c.Param("id")
	var doc models.CompanyDocument
	if err := config.DB.First(&doc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dokumen tidak ditemukan"})
		return
	}

	doc.Title = c.PostForm("title")
	doc.Category = c.PostForm("category")
	doc.DocDate = c.PostForm("doc_date")
	doc.Description = c.PostForm("description")

	file, err := c.FormFile("file")
	if err == nil {
		if file.Size > 10<<20 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "File terlalu besar. Maksimal 10 MB"})
			return
		}

		ext := filepath.Ext(file.Filename)
		allowedDocExts := map[string]bool{".pdf": true, ".doc": true, ".docx": true, ".xls": true, ".xlsx": true, ".ppt": true, ".pptx": true}
		if !allowedDocExts[ext] {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung. Gunakan PDF, DOC, XLS, atau PPT"})
			return
		}

		if doc.FileURL != "" {
			os.Remove("." + doc.FileURL)
		}

		doc.FileType = ext[1:]

		filename := strconv.FormatInt(time.Now().UnixNano(), 10) + ext
		savePath := filepath.Join("uploads", filename)

		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
			return
		}

		doc.FileURL = "/uploads/" + filename
		sizeKB := float64(file.Size) / 1024
		if sizeKB > 1024 {
			doc.FileSize = strconv.FormatFloat(sizeKB/1024, 'f', 1, 64) + " MB"
		} else {
			doc.FileSize = strconv.FormatFloat(sizeKB, 'f', 0, 64) + " KB"
		}
	}

	if err := config.DB.Save(&doc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengupdate data"})
		return
	}

	c.JSON(http.StatusOK, doc)
}

// DELETE /api/admin/company-documents/:id
func DeleteCompanyDocument(c *gin.Context) {
	id := c.Param("id")
	var doc models.CompanyDocument
	if err := config.DB.First(&doc, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dokumen tidak ditemukan"})
		return
	}

	if doc.FileURL != "" {
		os.Remove("." + doc.FileURL)
	}

	config.DB.Delete(&doc)
	c.JSON(http.StatusOK, gin.H{"message": "Dokumen berhasil dihapus"})
}
