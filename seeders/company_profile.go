package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedCompanyProfile() {
	var count int64
	config.DB.Model(&models.CompanyProfile{}).Count(&count)
	if count > 0 {
		return
	}

	profile := models.CompanyProfile{
		Title: "PT Tanjungenim Lestari Pulp and Paper",
		Content: `PT Tanjungenim Lestari Pulp and Paper (TeL) is world class manufacturer of high product quality and environmental friendly market pulp mill. This was established on June 18, 1990, commenced construction in mid-1997 and the commercial operation started on May, 2000. The mill is located in 1,250 ha area in the Banuayu village, District Empat Petulai Dangku, Muara Enim Regency, South Sumatra province, Indonesia.

TeL is a Foreign Investment Company (PMA)- Marubeni Corporation, Japan, as the National Vital Objects Industrial sector (OVNI) declared by the Minister of Industry in 2014. The mill has market pulp production capacity of 490,000 Adt / year. Presently mill has 1600 employees and support workforce together where ~ 80% of them are residents of South Sumatra`,
		Address: "",
		Phone:   "",
		Email:   "",
	}

	config.DB.Create(&profile)
	log.Println("Company profile seeded successfully")
}
