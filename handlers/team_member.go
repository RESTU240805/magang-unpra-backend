package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTeamMembers(c *gin.Context) {
	var members []models.TeamMember
	config.DB.Where("is_active = ?", true).Order("sort_order ASC, id ASC").Find(&members)
	c.JSON(http.StatusOK, gin.H{"data": members})
}

func GetAllTeamMembersAdmin(c *gin.Context) {
	var members []models.TeamMember
	config.DB.Order("sort_order ASC, id ASC").Find(&members)
	c.JSON(http.StatusOK, gin.H{"data": members})
}

type teamMemberInput struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	Description string `json:"description"`
	PhotoPath   string `json:"photo_path"`
	SortOrder   int    `json:"sort_order"`
	IsActive    bool   `json:"is_active"`
	IsFeatured  bool   `json:"is_featured"`
}

func CreateTeamMember(c *gin.Context) {
	var input teamMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	member := models.TeamMember{
		Name:        input.Name,
		Position:    input.Position,
		Description: input.Description,
		PhotoPath:   input.PhotoPath,
		SortOrder:   input.SortOrder,
		IsActive:    input.IsActive,
		IsFeatured:  input.IsFeatured,
	}
	config.DB.Create(&member)
	c.JSON(http.StatusCreated, gin.H{"data": member})
}

func UpdateTeamMember(c *gin.Context) {
	var member models.TeamMember
	if err := config.DB.First(&member, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
		return
	}

	var input teamMemberInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	member.Name = input.Name
	member.Position = input.Position
	member.Description = input.Description
	member.PhotoPath = input.PhotoPath
	member.SortOrder = input.SortOrder
	member.IsActive = input.IsActive
	member.IsFeatured = input.IsFeatured
	config.DB.Save(&member)
	c.JSON(http.StatusOK, gin.H{"data": member})
}

func DeleteTeamMember(c *gin.Context) {
	if err := config.DB.Delete(&models.TeamMember{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Team member deleted"})
}
