package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ── Public endpoint (tree) ──

func GetOrgStructure(c *gin.Context) {
	var groups []models.OrgGroup
	config.DB.Preload("Nodes", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Children", func(db2 *gorm.DB) *gorm.DB {
			return db2.Order("sort_order ASC, id ASC")
		}).Order("sort_order ASC, id ASC")
	}).Order("sort_order ASC, id ASC").Find(&groups)
	c.JSON(http.StatusOK, gin.H{"data": groups})
}

// ── OrgGroup CRUD ──

func GetAllOrgGroups(c *gin.Context) {
	var groups []models.OrgGroup
	config.DB.Preload("Nodes").Order("sort_order ASC, id ASC").Find(&groups)
	c.JSON(http.StatusOK, gin.H{"data": groups})
}

func CreateOrgGroup(c *gin.Context) {
	var input models.OrgGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateOrgGroup(c *gin.Context) {
	var group models.OrgGroup
	if err := config.DB.First(&group, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Group not found"})
		return
	}

	var input models.OrgGroup
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	group.Label = input.Label
	group.Color = input.Color
	group.SortOrder = input.SortOrder
	config.DB.Save(&group)
	c.JSON(http.StatusOK, gin.H{"data": group})
}

func DeleteOrgGroup(c *gin.Context) {
	id := c.Param("id")
	config.DB.Where("group_id = ?", id).Delete(&models.OrgNode{})
	if err := config.DB.Delete(&models.OrgGroup{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus grup"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Group deleted"})
}

// ── OrgNode CRUD ──

func GetAllOrgNodes(c *gin.Context) {
	var nodes []models.OrgNode
	config.DB.Order("sort_order ASC, id ASC").Find(&nodes)
	c.JSON(http.StatusOK, gin.H{"data": nodes})
}

func CreateOrgNode(c *gin.Context) {
	var input models.OrgNode
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateOrgNode(c *gin.Context) {
	var node models.OrgNode
	if err := config.DB.First(&node, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Node not found"})
		return
	}

	var input models.OrgNode
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	node.GroupID = input.GroupID
	node.ParentID = input.ParentID
	node.Name = input.Name
	node.Role = input.Role
	node.PhotoPath = input.PhotoPath
	node.SortOrder = input.SortOrder
	config.DB.Save(&node)
	c.JSON(http.StatusOK, gin.H{"data": node})
}

func DeleteOrgNode(c *gin.Context) {
	if err := config.DB.Delete(&models.OrgNode{}, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghapus posisi"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Node deleted"})
}

// ── Reset to template ──

func ResetOrgStructure(c *gin.Context) {
	config.DB.Exec("DELETE FROM org_nodes")
	config.DB.Exec("DELETE FROM org_groups")

	groups := []models.OrgGroup{
		{Label: "Dewan Komisaris", Color: "green", SortOrder: 1},
		{Label: "Dewan Direksi", Color: "blue", SortOrder: 2},
		{Label: "Kepala Divisi", Color: "green", SortOrder: 3},
	}
	for i := range groups {
		config.DB.Create(&groups[i])
	}

	nodes := []models.OrgNode{
		{GroupID: groups[0].ID, Name: "", Role: "Presiden Komisaris", SortOrder: 1},
		{GroupID: groups[0].ID, Name: "", Role: "Komisaris Independen", SortOrder: 2},
		{GroupID: groups[0].ID, Name: "", Role: "Komisaris", SortOrder: 3},
		{GroupID: groups[1].ID, Name: "", Role: "Presiden Direktur", SortOrder: 1},
		{GroupID: groups[1].ID, Name: "", Role: "Wakil Presiden Direktur", SortOrder: 2},
		{GroupID: groups[1].ID, Name: "", Role: "Direktur Keuangan", SortOrder: 3},
		{GroupID: groups[1].ID, Name: "", Role: "Direktur Operasional", SortOrder: 4},
		{GroupID: groups[1].ID, Name: "", Role: "Direktur Pemasaran", SortOrder: 5},
		{GroupID: groups[1].ID, Name: "", Role: "Direktur SDM", SortOrder: 6},
		{GroupID: groups[2].ID, Name: "", Role: "Kepala Divisi Keuangan", SortOrder: 1},
		{GroupID: groups[2].ID, Name: "", Role: "Kepala Divisi Operasional", SortOrder: 2},
		{GroupID: groups[2].ID, Name: "", Role: "Kepala Divisi Pemasaran", SortOrder: 3},
		{GroupID: groups[2].ID, Name: "", Role: "Kepala Divisi SDM / Umum", SortOrder: 4},
		{GroupID: groups[2].ID, Name: "", Role: "Kepala Divisi IT", SortOrder: 5},
	}
	for i := range nodes {
		config.DB.Create(&nodes[i])
	}

	c.JSON(http.StatusOK, gin.H{"message": "Struktur organisasi direset ke template awal"})
}
