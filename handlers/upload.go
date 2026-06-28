package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var allowedImageExts = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true,
}

var imageMimePrefixes = []string{"image/jpeg", "image/png", "image/gif", "image/webp"}

func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	if file.Size > 5<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File terlalu besar. Maksimal 5 MB"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedImageExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung. Gunakan jpg, jpeg, png, gif, atau webp"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
		return
	}
	defer src.Close()

	buf := make([]byte, 512)
	n, _ := src.Read(buf)
	mimeType := http.DetectContentType(buf[:n])
	mimeType = strings.ToLower(mimeType)
	validMime := false
	for _, prefix := range imageMimePrefixes {
		if strings.HasPrefix(mimeType, prefix) {
			validMime = true
			break
		}
	}
	if !validMime {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File bukan gambar valid"})
		return
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat folder"})
		return
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"url":      "/uploads/" + filename,
		"filename": filename,
	})
}

var allowedDocExts = map[string]bool{
	".pdf": true, ".doc": true, ".docx": true, ".xls": true, ".xlsx": true, ".ppt": true, ".pptx": true,
}

var docMimePrefixes = []string{"application/pdf", "application/msword", "application/vnd.openxmlformats-officedocument", "application/vnd.ms-excel", "application/vnd.ms-powerpoint"}

func UploadDocument(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File tidak ditemukan"})
		return
	}

	if file.Size > 20<<20 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File terlalu besar. Maksimal 20 MB"})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedDocExts[ext] {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Format file tidak didukung. Gunakan pdf, doc, docx, xls, xlsx, ppt, atau pptx"})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membaca file"})
		return
	}
	defer src.Close()

	buf := make([]byte, 512)
	n, _ := src.Read(buf)
	mimeType := http.DetectContentType(buf[:n])
	mimeType = strings.ToLower(mimeType)
	validMime := false
	for _, prefix := range docMimePrefixes {
		if strings.HasPrefix(mimeType, prefix) {
			validMime = true
			break
		}
	}
	if !validMime {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File dokumen tidak valid"})
		return
	}

	uploadDir := "./uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat folder"})
		return
	}

	filename := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	filePath := filepath.Join(uploadDir, filename)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan file"})
		return
	}

	fileURL := fmt.Sprintf("uploads/%s", filename)
	c.JSON(http.StatusOK, gin.H{
		"url":      fileURL,
		"filename": filename,
	})
}
