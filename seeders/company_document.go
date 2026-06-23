package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedCompanyDocuments() {
	var count int64
	config.DB.Model(&models.CompanyDocument{}).Count(&count)
	if count > 0 {
		return
	}

	docs := []models.CompanyDocument{
		{
			Title:       "PEFC Chain of Custody Commitment Statement",
			Category:    "DOCUMENT PREVIEW",
			DocDate:     "April 2023",
			FileType:    "PDF",
			FileSize:    "1.2 MB",
			Description: "PT. Tanjungenim Lestari Pulp and Paper is committed to sourcing wood-based products in an environmentally and socially responsible manner, in full compliance with EU Timber Regulation and sustainable forest management practices.",
			FileURL:     "/files/PEFC_Statement.pdf",
		},
	}

	for _, d := range docs {
		config.DB.Create(&d)
	}
	log.Println("Company documents seeded successfully")
}
