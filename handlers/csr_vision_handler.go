package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ─── CSR Vision Content ───────────────────────────────────

const defaultCsrCorporateDescription = `Corporate Social Responsibility is defined as sustainable commitment to business to act ethically, operate legally and contribute to improving the economy while at the same time improving the quality of life of employees and their families, local communities and society as a whole around the company.

The CSR program is an investment for companies for the growth and sustainability of the company and is no longer considered as a cost center, but as a profit center. The CSR program is the company's commitment to support the creation of sustainable development.

The implementation of CSR programs is one form of implementation of the concept of Good Corporate Governance as a social and environmental responsibility as stipulated in the regulations / Law no. 40 of 2007 concerning Limited Liability Companies.`

const defaultCsrObjectives = `The purpose of implementing the Corporate Social Responsibility (CSR) of PT. Tanjungenim Lestari Pulp and Paper is not limited to fulfilling the Company's responsibilities as well as the form of compliance with regulations only. Furthermore, a well-implemented and systematic CSR program can form a harmonious, balanced relationship and support between the Company and the community in the surrounding.

The company also believes that CSR programs are an important part that is inseparable from the benchmark of business success, apart from the implementation of good management and operational performance.

On this basis, through the implementation of the responsibilities contained in the vision, mission and corporate values, the Company outlines CSR policies and programs that meet sustainability standards, have integrity, uphold business ethics, and comply with laws and regulations that apply. These standards are expected to create CSR programs that can have a positive, effective and targeted impact, be able to empower the community's capabilities, and be sustainable in the long term.`

type csrVisionContentInput struct {
	CorporateDescription *string `json:"corporate_description"`
	Objectives           *string `json:"objectives"`
}

func getOrCreateCsrVisionContent() models.CsrVisionContent {
	var item models.CsrVisionContent
	if err := config.DB.First(&item).Error; err != nil {
		item = models.CsrVisionContent{
			CorporateDescription: defaultCsrCorporateDescription,
			Objectives:           defaultCsrObjectives,
		}
		config.DB.Create(&item)
	} else if item.Objectives == "" {
		item.Objectives = defaultCsrObjectives
		config.DB.Save(&item)
	}
	return item
}

func GetCsrVisionContent(c *gin.Context) {
	item := getOrCreateCsrVisionContent()
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func UpdateCsrVisionContent(c *gin.Context) {
	item := getOrCreateCsrVisionContent()
	var input csrVisionContentInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	if input.CorporateDescription != nil {
		item.CorporateDescription = *input.CorporateDescription
	}
	if input.Objectives != nil {
		item.Objectives = *input.Objectives
	}
	config.DB.Save(&item)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

// ─── CSR Vision Strategies ────────────────────────────────

func GetAllCsrVisionStrategies(c *gin.Context) {
	var items []models.CsrVisionStrategy
	config.DB.Where("is_active = ?", true).Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

func GetAllCsrVisionStrategiesAdmin(c *gin.Context) {
	var items []models.CsrVisionStrategy
	config.DB.Order("sort_order asc").Find(&items)
	c.JSON(http.StatusOK, gin.H{"data": items})
}

type csrVisionStrategyInput struct {
	Description *string `json:"description"`
	SortOrder   *int    `json:"sort_order"`
	IsActive    *bool   `json:"is_active"`
}

func CreateCsrVisionStrategy(c *gin.Context) {
	var input csrVisionStrategyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	item := models.CsrVisionStrategy{
		IsActive: true,
	}
	if input.Description != nil {
		item.Description = *input.Description
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

func UpdateCsrVisionStrategy(c *gin.Context) {
	id := c.Param("id")
	var input csrVisionStrategyInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	updates := map[string]interface{}{}
	if input.Description != nil {
		updates["description"] = *input.Description
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
	result := config.DB.Model(&models.CsrVisionStrategy{}).Where("id = ?", id).Updates(updates)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Strategy not found"})
		return
	}
	var item models.CsrVisionStrategy
	config.DB.First(&item, id)
	c.JSON(http.StatusOK, gin.H{"data": item})
}

func DeleteCsrVisionStrategy(c *gin.Context) {
	config.DB.Delete(&models.CsrVisionStrategy{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Strategy deleted"})
}
