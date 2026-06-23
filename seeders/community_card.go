package seeders

import (
	"log"
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
)

func SeedCommunityCards() {
	var count int64
	config.DB.Model(&models.CommunityCard{}).Count(&count)

	cards := []models.CommunityCard{
		{
			Title:       "Education",
			Description: "Prepare a generation of quality sources from internal employees as well as from the village community around the company of Telpp. We achieved this through scholarships, traineeships, Development of Community Education, School foundation program and school facilities & infrastructure improvement.",
			IconPath:    "uploads/community-edu.jpg",
			Link:        "/sustainability/csr/community",
			SortOrder:   0,
			IsActive:    true,
		},
		{
			Title:       "Infrastructure",
			Description: "Actively participating to assist availability of public facilities as the driving force of economic growth and improving the quality of life and welfare of local communities in the village around the company PT Tanjungenim Lestari Pulp and Paper.",
			IconPath:    "uploads/community-infra.png",
			Link:        "/sustainability/csr/community",
			SortOrder:   1,
			IsActive:    true,
		},
		{
			Title:       "Local Economy Development",
			Description: "Our beneficiary community businesses of Small, Medium Enterprise Development that have an impact on the direct, indirect and the ability of people involved in order to become economically independent and sustainable primarily around the company of PT Tanjungenim Lestari Pulp and Paper.",
			IconPath:    "uploads/community-economy.jpeg",
			Link:        "/sustainability/csr/community",
			SortOrder:   2,
			IsActive:    true,
		},
		{
			Title:       "Health & Conservation",
			Description: "Increasing community awareness, changing the mindset of not stakeholders and the community in health development by increasing community empowerment efforts through quality appropriate service facilities and community movements in healthy living.",
			IconPath:    "uploads/community-health.png",
			Link:        "/sustainability/csr/community",
			SortOrder:   3,
			IsActive:    true,
		},
	}

	for _, c := range cards {
		var existing models.CommunityCard
		result := config.DB.Where("title = ?", c.Title).First(&existing)
		if result.Error != nil {
			config.DB.Create(&c)
		} else {
			existing.IconPath = c.IconPath
			config.DB.Save(&existing)
		}
	}
	log.Println("Community cards seeded successfully")
}
