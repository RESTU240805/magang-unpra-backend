package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedCreeds() {
	var count int64
	config.DB.Model(&models.Creed{}).Count(&count)
	if count > 0 {
		return
	}

	creeds := []models.Creed{
		{
			TitleJP:     "和",
			TitleEN:     "Harmony",
			Roma:        "WA",
			Tagline:     "To respect each other and cooperate.",
			Description: "We shall stay in touch with society and stakeholders by engaging in corporate activities which advance our credibility as preferred principals.",
			SortOrder:   0,
			IsActive:    true,
		},
		{
			TitleJP:     "新",
			TitleEN:     "Innovation",
			Roma:        "SHIN",
			Tagline:     "To be active and innovative.",
			Description: "We shall constantly strive hard to improve our performance.",
			SortOrder:   1,
			IsActive:    true,
		},
		{
			TitleJP:     "心",
			TitleEN:     "Heart",
			Roma:        "SHIN",
			Tagline:     "To be fair and decent.",
			Description: "We shall comply with the laws and follow fair corporate practices.",
			SortOrder:   2,
			IsActive:    true,
		},
	}

	for _, c := range creeds {
		config.DB.Create(&c)
	}
	log.Println("Creeds seeded successfully")
}
