package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ─── ContactInfo ──────────────────────────────────────────────

func GetContactInfo(c *gin.Context) {
	var info models.ContactInfo
	config.DB.First(&info)
	c.JSON(http.StatusOK, gin.H{"data": info})
}

type contactInfoInput struct {
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Address      string `json:"address"`
	WorkingHours string `json:"working_hours"`
	Copyright    string `json:"copyright"`
	HeroImage    string `json:"hero_image"`
	MapLink      string `json:"map_link"`
}

func UpdateContactInfo(c *gin.Context) {
	var input contactInfoInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	var info models.ContactInfo
	config.DB.First(&info)
	info.Email = input.Email
	info.Phone = input.Phone
	info.Address = input.Address
	info.WorkingHours = input.WorkingHours
	info.Copyright = input.Copyright
	info.HeroImage = input.HeroImage
	info.MapLink = input.MapLink
	config.DB.Save(&info)
	c.JSON(http.StatusOK, gin.H{"data": info})
}

// ─── ContactOffice ────────────────────────────────────────────

func GetAllContactOffices(c *gin.Context) {
	var offices []models.ContactOffice
	config.DB.Order("sort_order asc").Find(&offices)
	c.JSON(http.StatusOK, gin.H{"data": offices})
}

func GetAllContactOfficesAdmin(c *gin.Context) {
	var offices []models.ContactOffice
	config.DB.Order("sort_order asc").Find(&offices)
	c.JSON(http.StatusOK, gin.H{"data": offices})
}

type contactOfficeInput struct {
	Name      string `json:"name"`
	Address   string `json:"address"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	MapLink   string `json:"map_link"`
	Image     string `json:"image"`
	SortOrder int    `json:"sort_order"`
	IsActive  *bool  `json:"is_active"`
}

func CreateContactOffice(c *gin.Context) {
	var input contactOfficeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	office := models.ContactOffice{
		Name:      input.Name,
		Address:   input.Address,
		Phone:     input.Phone,
		Email:     input.Email,
		MapLink:   input.MapLink,
		Image:     input.Image,
		SortOrder: input.SortOrder,
		IsActive:  true,
	}
	if input.IsActive != nil {
		office.IsActive = *input.IsActive
	}
	config.DB.Create(&office)
	c.JSON(http.StatusCreated, gin.H{"data": office})
}

func UpdateContactOffice(c *gin.Context) {
	var office models.ContactOffice
	id := c.Param("id")
	if err := config.DB.First(&office, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Office not found"})
		return
	}
	var input contactOfficeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	office.Name = input.Name
	office.Address = input.Address
	office.Phone = input.Phone
	office.Email = input.Email
	office.MapLink = input.MapLink
	office.Image = input.Image
	office.SortOrder = input.SortOrder
	if input.IsActive != nil {
		office.IsActive = *input.IsActive
	}
	config.DB.Save(&office)
	c.JSON(http.StatusOK, gin.H{"data": office})
}

func DeleteContactOffice(c *gin.Context) {
	config.DB.Delete(&models.ContactOffice{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Office deleted"})
}
