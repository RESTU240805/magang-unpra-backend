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

func GetTeamMemberById(c *gin.Context) {
	var member models.TeamMember
	if err := config.DB.First(&member, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": member})
}

func CreateTeamMember(c *gin.Context) {
	var input models.TeamMember
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateTeamMember(c *gin.Context) {
	var member models.TeamMember
	if err := config.DB.First(&member, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Team member not found"})
		return
	}

	var input models.TeamMember
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
