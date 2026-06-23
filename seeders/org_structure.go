package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedOrgStructure() {
	var count int64
	config.DB.Model(&models.OrgGroup{}).Count(&count)
	if count > 0 {
		return
	}

	config.DB.Where("1 = 1").Delete(&models.OrgNode{})
	config.DB.Where("1 = 1").Delete(&models.OrgGroup{})

	groups := []models.OrgGroup{
		{Label: "Dewan Komisaris", Color: "green", SortOrder: 1},
		{Label: "Dewan Direksi", Color: "blue", SortOrder: 2},
		{Label: "Kepala Divisi", Color: "green", SortOrder: 3},
	}

	for i := range groups {
		config.DB.Create(&groups[i])
	}

	nodes := []models.OrgNode{
		// ── Level 1: Dewan Komisaris (top) ──
		{GroupID: groups[0].ID, ParentID: nil, Name: "", Role: "Presiden Komisaris", SortOrder: 1},
		{GroupID: groups[0].ID, ParentID: nil, Name: "", Role: "Komisaris Independen", SortOrder: 2},
		{GroupID: groups[0].ID, ParentID: nil, Name: "", Role: "Komisaris", SortOrder: 3},

		// ── Level 2: Dewan Direksi (middle) ──
		{GroupID: groups[1].ID, ParentID: nil, Name: "", Role: "Presiden Direktur", SortOrder: 1},
		{GroupID: groups[1].ID, ParentID: nil, Name: "", Role: "Wakil Presiden Direktur", SortOrder: 2},
		{GroupID: groups[1].ID, ParentID: nil, Name: "", Role: "Direktur Keuangan", SortOrder: 3},
		{GroupID: groups[1].ID, ParentID: nil, Name: "", Role: "Direktur Operasional", SortOrder: 4},
		{GroupID: groups[1].ID, ParentID: nil, Name: "", Role: "Direktur Pemasaran", SortOrder: 5},
		{GroupID: groups[1].ID, ParentID: nil, Name: "", Role: "Direktur SDM", SortOrder: 6},

		// ── Level 3: Kepala Divisi (bottom) ──
		{GroupID: groups[2].ID, ParentID: nil, Name: "", Role: "Kepala Divisi Keuangan", SortOrder: 1},
		{GroupID: groups[2].ID, ParentID: nil, Name: "", Role: "Kepala Divisi Operasional", SortOrder: 2},
		{GroupID: groups[2].ID, ParentID: nil, Name: "", Role: "Kepala Divisi Pemasaran", SortOrder: 3},
		{GroupID: groups[2].ID, ParentID: nil, Name: "", Role: "Kepala Divisi SDM / Umum", SortOrder: 4},
		{GroupID: groups[2].ID, ParentID: nil, Name: "", Role: "Kepala Divisi IT", SortOrder: 5},
	}

	for i := range nodes {
		config.DB.Create(&nodes[i])
	}

	log.Println("Org structure template seeded!")
}
