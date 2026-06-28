package handlers

import (
	"magang-unpra-backend/config"
	"magang-unpra-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPeopleDevelopmentPage(c *gin.Context) {
	var page models.PeopleDevelopmentPage
	if err := config.DB.First(&page).Error; err != nil {
		page = models.PeopleDevelopmentPage{Description: defaultPeopleDevelopmentDescription}
		config.DB.Create(&page)
	}
	c.JSON(http.StatusOK, gin.H{"data": page})
}

func UpdatePeopleDevelopmentPage(c *gin.Context) {
	var page models.PeopleDevelopmentPage
	if err := config.DB.First(&page).Error; err != nil {
		page = models.PeopleDevelopmentPage{}
		config.DB.Create(&page)
	}

	var input struct {
		Description string `json:"description"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	page.Description = input.Description
	config.DB.Save(&page)
	c.JSON(http.StatusOK, gin.H{"data": page})
}

func GetAllPeopleDevelopmentPillars(c *gin.Context) {
	var pillars []models.PeopleDevelopmentPillar
	config.DB.Order("sort_order asc").Find(&pillars)
	if len(pillars) == 0 {
		pillars = seedDefaultPeopleDevelopmentPillars()
	}
	c.JSON(http.StatusOK, gin.H{"data": pillars})
}

type peopleDevelopmentPillarInput struct {
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	SortOrder int    `json:"sort_order"`
}

func CreatePeopleDevelopmentPillar(c *gin.Context) {
	var input peopleDevelopmentPillarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	pillar := models.PeopleDevelopmentPillar{Title: input.Title, Desc: input.Desc, SortOrder: input.SortOrder}
	config.DB.Create(&pillar)
	c.JSON(http.StatusCreated, gin.H{"data": pillar})
}

func UpdatePeopleDevelopmentPillar(c *gin.Context) {
	var pillar models.PeopleDevelopmentPillar
	if err := config.DB.First(&pillar, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pillar not found"})
		return
	}

	var input peopleDevelopmentPillarInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	pillar.Title = input.Title
	pillar.Desc = input.Desc
	pillar.SortOrder = input.SortOrder
	config.DB.Save(&pillar)
	c.JSON(http.StatusOK, gin.H{"data": pillar})
}

func DeletePeopleDevelopmentPillar(c *gin.Context) {
	config.DB.Delete(&models.PeopleDevelopmentPillar{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Pillar deleted"})
}

func GetAllPeopleDevelopmentSliders(c *gin.Context) {
	var slides []models.PeopleDevelopmentSlider
	config.DB.Order("sort_order asc").Find(&slides)
	if len(slides) == 0 {
		slides = seedDefaultPeopleDevelopmentSliders()
	}
	c.JSON(http.StatusOK, gin.H{"data": slides})
}

type peopleDevelopmentSliderInput struct {
	ImageURL  string `json:"image_url"`
	Caption   string `json:"caption"`
	SortOrder int    `json:"sort_order"`
}

func CreatePeopleDevelopmentSlider(c *gin.Context) {
	var input peopleDevelopmentSliderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}
	slide := models.PeopleDevelopmentSlider{ImageURL: input.ImageURL, Caption: input.Caption, SortOrder: input.SortOrder}
	config.DB.Create(&slide)
	c.JSON(http.StatusCreated, gin.H{"data": slide})
}

func UpdatePeopleDevelopmentSlider(c *gin.Context) {
	var slide models.PeopleDevelopmentSlider
	if err := config.DB.First(&slide, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Slider not found"})
		return
	}

	var input peopleDevelopmentSliderInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Input tidak valid"})
		return
	}

	slide.ImageURL = input.ImageURL
	slide.Caption = input.Caption
	slide.SortOrder = input.SortOrder
	config.DB.Save(&slide)
	c.JSON(http.StatusOK, gin.H{"data": slide})
}

func DeletePeopleDevelopmentSlider(c *gin.Context) {
	config.DB.Delete(&models.PeopleDevelopmentSlider{}, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"message": "Slider deleted"})
}

const defaultPeopleDevelopmentDescription = `<p>People development is carried out in order to prepare the best talent to be able to show the best performance and to achieve the business targets they carry. Talking about business is not only about running but also how they manage the business even more and sustain it in the future.</p><p>PT Tanjungenim Lestari Pulp and Paper are committed to developing competent, motivated and integrated information systems to achieve a competitive organization. The commitment is realized in several activities and achievements as follows:</p>`

func seedDefaultPeopleDevelopmentPillars() []models.PeopleDevelopmentPillar {
	pillars := []models.PeopleDevelopmentPillar{
		{Title: "Skills-Based Leadership Coaching", Desc: "Achieve a high performing organization by developing skills-based Coaching for leadership styles through a comprehensive Supervisory / Management / Leadership Program.", SortOrder: 1},
		{Title: "Succession Planning & Career Path", Desc: "Develop and implement succession plans for critical positions to ensure the regeneration of employees through the Career Path and Talent Management.", SortOrder: 2},
		{Title: "Community & Scholarship Programs", Desc: "Prepare a generation of quality sources from internal employees as well as from the village community through scholarships, internship, community education, school foundation programs and infrastructure improvement.", SortOrder: 3},
		{Title: "Technology & Information Skills", Desc: "Improve knowledge and skills related to technology and information across all levels of the organization.", SortOrder: 4},
		{Title: "Rules, Competencies & Knowledge", Desc: "Meet the rules and competencies with good knowledge and skills to maintain professional standards throughout the company.", SortOrder: 5},
	}
	for i := range pillars {
		config.DB.Create(&pillars[i])
	}
	return pillars
}

func seedDefaultPeopleDevelopmentSliders() []models.PeopleDevelopmentSlider {
	slides := []models.PeopleDevelopmentSlider{
		{ImageURL: "/images/Training-First-Aid.jpg", Caption: "First Aid Training Session", SortOrder: 1},
		{ImageURL: "/images/Management-Safety-Patrol.jpg", Caption: "Safety Patrol Management", SortOrder: 2},
		{ImageURL: "/images/Industrial-Hygiene.jpg", Caption: "Team Safety Workshop", SortOrder: 3},
		{ImageURL: "/images/people1.jpeg", Caption: "Industrial Hygiene Program", SortOrder: 4},
		{ImageURL: "/images/people2.jpeg", Caption: "People Development Initiative", SortOrder: 5},
	}
	for i := range slides {
		config.DB.Create(&slides[i])
	}
	return slides
}
