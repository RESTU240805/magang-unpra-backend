package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ─── CSR Reports ──────────────────────────────────────────

func GetAllCsrReports(c *gin.Context) {
	var items []models.CsrReport
	config.DB.Where("is_active = ?", true).Order("year desc, sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllCsrReportsAdmin(c *gin.Context) {
	var items []models.CsrReport
	config.DB.Order("year desc, sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type csrReportInput struct {
	Year      *int    `json:"year"`
	Quarter   *string `json:"quarter"`
	Period    *string `json:"period"`
	FileURL   *string `json:"file_url"`
	SortOrder *int    `json:"sort_order"`
	IsActive  *bool   `json:"is_active"`
}

func CreateCsrReport(c *gin.Context) {
	var input csrReportInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.CsrReport{
		IsActive: true,
	}
	if input.Year != nil {
		item.Year = *input.Year
	}
	if input.Quarter != nil {
		item.Quarter = *input.Quarter
	}
	if input.Period != nil {
		item.Period = *input.Period
	}
	if input.FileURL != nil {
		item.FileURL = *input.FileURL
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

func UpdateCsrReport(c *gin.Context) {
	id := c.Param("id")
	var input csrReportInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	updates := map[string]interface{}{}
	if input.Year != nil {
		updates["year"] = *input.Year
	}
	if input.Quarter != nil {
		updates["quarter"] = *input.Quarter
	}
	if input.Period != nil {
		updates["period"] = *input.Period
	}
	if input.FileURL != nil {
		updates["file_url"] = *input.FileURL
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
	result := config.DB.Model(&models.CsrReport{}).Where("id = ?", id).Updates(updates)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}
	var item models.CsrReport
	config.DB.First(&item, id)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteCsrReport(c *gin.Context) {
	config.DB.Delete(&models.CsrReport{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Report deleted"})
}
