package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedMenus() {
	var count int64
	config.DB.Model(&models.Menu{}).Count(&count)
	if count > 0 {
		return
	}

	type menuDef struct {
		Name     string
		URL      string
		Children []menuDef
	}

	defs := []menuDef{
		{Name: "HOME", URL: "/"},
		{Name: "OUR COMPANY", URL: "", Children: []menuDef{
			{Name: "Our Company", URL: "/our-company"},
			{Name: "Our Team", URL: "/our-team"},
		}},
		{Name: "PRODUCT", URL: "/product"},
		{Name: "NEWS", URL: "/news"},
		{Name: "SUSTAINABILITY", URL: "", Children: []menuDef{
			{Name: "Forest Management", URL: "/sustainability/forest-management"},
			{Name: "People Development", URL: "/sustainability/people-development"},
			{Name: "Supply Chain", URL: "/sustainability/supply-chain"},
			{Name: "Pulp Process", URL: "/sustainability/pulp-process"},
			{Name: "Safety & Health", URL: "/sustainability/safety-health"},
			{Name: "Corporate Social Responsibility", URL: "", Children: []menuDef{
				{Name: "Vision And Mission", URL: "/sustainability/csr/vision"},
				{Name: "Local Community Development", URL: "/sustainability/csr/community"},
				{Name: "CSR Report", URL: "/sustainability/csr/report"},
			}},
		}},
		{Name: "BIODIVERSITY", URL: "https://kehati.telpp.com/"},
		{Name: "E-PROC", URL: "https://eproc.telpp.com/_gst_home.php"},
		{Name: "E-RECRUITMENT", URL: "https://erecruitment.telpp.com/er_tlpp/"},
		{Name: "CONTACT", URL: "/contact"},
	}

	for i, d := range defs {
		parent := models.Menu{
			Name:      d.Name,
			URL:       d.URL,
			IsActive:  true,
			SortOrder: i,
		}
		config.DB.Create(&parent)
		for j, child := range d.Children {
			childMenu := models.Menu{
				Name:      child.Name,
				URL:       child.URL,
				IsActive:  true,
				ParentID:  &parent.ID,
				SortOrder: j,
			}
			config.DB.Create(&childMenu)
			for k, grandchild := range child.Children {
				gcMenu := models.Menu{
					Name:      grandchild.Name,
					URL:       grandchild.URL,
					IsActive:  true,
					ParentID:  &childMenu.ID,
					SortOrder: k,
				}
				config.DB.Create(&gcMenu)
			}
		}
	}

	log.Println("Menus seeded successfully")
}
