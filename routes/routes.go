package routes

import (
	"os"
	"strings"
	"time"

	"magang-unpra-backend/handlers"
	"magang-unpra-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.SecurityHeaders())

	allowOrigins := os.Getenv("CORS_ORIGINS")
	if allowOrigins == "" {
		allowOrigins = "http://localhost:5173"
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     strings.Split(allowOrigins, ","),
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve public static files
	r.Static("/uploads", "./uploads")
	r.Static("/pulp-process", "./pulp proces")
	r.Static("/safety", "./safety")
	r.Static("/images", "../magang-unpra-frontend/public/images")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	login := r.Group("/api/login")
	login.Use(middleware.RateLimit(5, time.Minute))
	login.POST("", handlers.Login)

	upload := r.Group("/api/upload")
	upload.Use(middleware.AuthRequired())
	upload.Use(middleware.RateLimit(10, time.Minute))
		upload.POST("", handlers.UploadImage)
		upload.POST("/document", handlers.UploadDocument)

	api := r.Group("/api")
	api.Use(middleware.RateLimit(60, time.Minute))
	{
		// Auth
		api.GET("/auth/me", middleware.AuthRequired(), handlers.GetCurrentUser)

		// Public
		api.GET("/news", handlers.GetAllNews)
		api.GET("/news/:id", handlers.GetNewsById)
		api.GET("/products", handlers.GetAllProducts)
		api.GET("/products/:id", handlers.GetProductById)
		api.GET("/company-profile", handlers.GetCompanyProfile)
		api.GET("/about-section", handlers.GetAboutSection)
		api.GET("/product-slides", handlers.GetAllSlides)
		api.GET("/product-page", handlers.GetProductPage)
		api.GET("/team-members", handlers.GetAllTeamMembers)
		api.GET("/org-chart", handlers.GetOrgChart)
		api.GET("/community-cards", handlers.GetAllCommunityCards)
		api.GET("/creeds", handlers.GetAllCreeds)
		api.GET("/company-documents", handlers.GetAllCompanyDocuments)
		api.GET("/forest-wood-types", handlers.GetAllWoodTypes)
		api.GET("/forest-approach", handlers.GetForestApproach)
		api.GET("/forest-sliders", handlers.GetAllForestSliders)
		api.GET("/people-development-page", handlers.GetPeopleDevelopmentPage)
		api.GET("/people-development-pillars", handlers.GetAllPeopleDevelopmentPillars)
		api.GET("/people-development-sliders", handlers.GetAllPeopleDevelopmentSliders)

		// ─── New Sustainability Public Routes ───────────
		api.GET("/pulp-process-sections", handlers.GetAllPulpProcessSections)
		api.GET("/pulp-process-recoveries", handlers.GetAllPulpProcessRecoveries)
		api.GET("/safety-policies", handlers.GetAllSafetyPolicies)
		api.GET("/safety-k3-targets", handlers.GetAllSafetyK3Targets)
		api.GET("/safety-k3-programs", handlers.GetAllSafetyK3Programs)
		api.GET("/safety-sliders", handlers.GetAllSafetySliders)
		api.GET("/supply-chain-strategies", handlers.GetAllSupplyChainStrategies)
		api.GET("/supply-chain-sustainability-items", handlers.GetAllSupplyChainSustainabilityItems)
		api.GET("/supply-chain-policies", handlers.GetAllSupplyChainPolicies)
		api.GET("/csr-vision-content", handlers.GetCsrVisionContent)
		api.GET("/csr-vision-strategies", handlers.GetAllCsrVisionStrategies)
		api.GET("/csr-reports", handlers.GetAllCsrReports)
		api.GET("/contact-info", handlers.GetContactInfo)
		api.GET("/contact-offices", handlers.GetAllContactOffices)
		api.GET("/menus", handlers.GetActiveMenus)

		// Admin
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired())
		admin.Use(middleware.RoleRequired("admin"))
		{
			admin.GET("/news", handlers.GetAllNewsAdmin)
			admin.POST("/news", handlers.CreateNews)
			admin.PUT("/news/:id", handlers.UpdateNews)
			admin.DELETE("/news/:id", handlers.DeleteNews)

			admin.GET("/products", handlers.GetAllProducts)
			admin.POST("/products", handlers.CreateProduct)
			admin.PUT("/products/:id", handlers.UpdateProduct)
			admin.DELETE("/products/:id", handlers.DeleteProduct)

			admin.GET("/product-slides", handlers.GetAllSlides)
			admin.POST("/product-slides", handlers.CreateSlide)
			admin.PUT("/product-slides/:id", handlers.UpdateSlide)
			admin.DELETE("/product-slides/:id", handlers.DeleteSlide)

			admin.GET("/company-profile", handlers.GetCompanyProfile)
			admin.PUT("/company-profile", handlers.UpdateCompanyProfile)
			admin.GET("/about-section", handlers.GetAboutSection)
			admin.PUT("/about-section", handlers.UpdateAboutSection)
			admin.GET("/product-page", handlers.GetProductPage)
			admin.PUT("/product-page", handlers.UpdateProductPage)

			admin.GET("/org-structure", handlers.GetOrgStructure)
			admin.GET("/org-chart", handlers.GetOrgChart)

			admin.GET("/team-members", handlers.GetAllTeamMembersAdmin)
			admin.POST("/team-members", handlers.CreateTeamMember)
			admin.PUT("/team-members/:id", handlers.UpdateTeamMember)
			admin.DELETE("/team-members/:id", handlers.DeleteTeamMember)

			admin.GET("/community-cards", handlers.GetAllCommunityCardsAdmin)
			admin.POST("/community-cards", handlers.CreateCommunityCard)
			admin.PUT("/community-cards/:id", handlers.UpdateCommunityCard)
			admin.DELETE("/community-cards/:id", handlers.DeleteCommunityCard)

			admin.GET("/org-groups", handlers.GetAllOrgGroups)
			admin.POST("/org-groups", handlers.CreateOrgGroup)
			admin.PUT("/org-groups/:id", handlers.UpdateOrgGroup)
			admin.DELETE("/org-groups/:id", handlers.DeleteOrgGroup)

			admin.GET("/org-nodes", handlers.GetAllOrgNodes)
			admin.POST("/org-nodes", handlers.CreateOrgNode)
			admin.PUT("/org-nodes/:id", handlers.UpdateOrgNode)
			admin.DELETE("/org-nodes/:id", handlers.DeleteOrgNode)
			admin.POST("/org-structure/reset", handlers.ResetOrgStructure)
			admin.PUT("/org-chart", handlers.UpdateOrgChart)
			// Tambahan baru - Creed
			admin.GET("/creeds", handlers.GetAllCreeds)
			admin.POST("/creeds", handlers.CreateCreed)
			admin.PUT("/creeds/:id", handlers.UpdateCreed)
			admin.DELETE("/creeds/:id", handlers.DeleteCreed)

			// Tambahan baru - Company Document
			admin.GET("/company-documents", handlers.GetAllCompanyDocuments)
			admin.POST("/company-documents", handlers.CreateCompanyDocument)
			admin.PUT("/company-documents/:id", handlers.UpdateCompanyDocument)
			admin.DELETE("/company-documents/:id", handlers.DeleteCompanyDocument)

			admin.GET("/forest-wood-types", handlers.GetAllWoodTypes)
			admin.POST("/forest-wood-types", handlers.CreateWoodType)
			admin.PUT("/forest-wood-types/:id", handlers.UpdateWoodType)
			admin.DELETE("/forest-wood-types/:id", handlers.DeleteWoodType)

			admin.GET("/forest-approach", handlers.GetForestApproach)
			admin.PUT("/forest-approach", handlers.UpdateForestApproach)

			admin.GET("/forest-sliders", handlers.GetAllForestSliders)
			admin.POST("/forest-sliders", handlers.CreateForestSlider)
			admin.PUT("/forest-sliders/:id", handlers.UpdateForestSlider)
			admin.DELETE("/forest-sliders/:id", handlers.DeleteForestSlider)

			admin.GET("/people-development-page", handlers.GetPeopleDevelopmentPage)
			admin.PUT("/people-development-page", handlers.UpdatePeopleDevelopmentPage)
			admin.GET("/people-development-pillars", handlers.GetAllPeopleDevelopmentPillars)
			admin.POST("/people-development-pillars", handlers.CreatePeopleDevelopmentPillar)
			admin.PUT("/people-development-pillars/:id", handlers.UpdatePeopleDevelopmentPillar)
			admin.DELETE("/people-development-pillars/:id", handlers.DeletePeopleDevelopmentPillar)
			admin.GET("/people-development-sliders", handlers.GetAllPeopleDevelopmentSliders)
			admin.POST("/people-development-sliders", handlers.CreatePeopleDevelopmentSlider)
			admin.PUT("/people-development-sliders/:id", handlers.UpdatePeopleDevelopmentSlider)
			admin.DELETE("/people-development-sliders/:id", handlers.DeletePeopleDevelopmentSlider)

			// ─── Pulp Process Admin ─────────────────────
			admin.GET("/pulp-process-sections", handlers.GetAllPulpProcessSectionsAdmin)
			admin.POST("/pulp-process-sections", handlers.CreatePulpProcessSection)
			admin.PUT("/pulp-process-sections/:id", handlers.UpdatePulpProcessSection)
			admin.DELETE("/pulp-process-sections/:id", handlers.DeletePulpProcessSection)

			admin.GET("/pulp-process-recoveries", handlers.GetAllPulpProcessRecoveriesAdmin)
			admin.POST("/pulp-process-recoveries", handlers.CreatePulpProcessRecovery)
			admin.PUT("/pulp-process-recoveries/:id", handlers.UpdatePulpProcessRecovery)
			admin.DELETE("/pulp-process-recoveries/:id", handlers.DeletePulpProcessRecovery)

			// ─── Safety & Health Admin ──────────────────
			admin.GET("/safety-policies", handlers.GetAllSafetyPoliciesAdmin)
			admin.POST("/safety-policies", handlers.CreateSafetyPolicy)
			admin.PUT("/safety-policies/:id", handlers.UpdateSafetyPolicy)
			admin.DELETE("/safety-policies/:id", handlers.DeleteSafetyPolicy)

			admin.GET("/safety-k3-targets", handlers.GetAllSafetyK3TargetsAdmin)
			admin.POST("/safety-k3-targets", handlers.CreateSafetyK3Target)
			admin.PUT("/safety-k3-targets/:id", handlers.UpdateSafetyK3Target)
			admin.DELETE("/safety-k3-targets/:id", handlers.DeleteSafetyK3Target)

			admin.GET("/safety-k3-programs", handlers.GetAllSafetyK3ProgramsAdmin)
			admin.POST("/safety-k3-programs", handlers.CreateSafetyK3Program)
			admin.PUT("/safety-k3-programs/:id", handlers.UpdateSafetyK3Program)
			admin.DELETE("/safety-k3-programs/:id", handlers.DeleteSafetyK3Program)

			admin.GET("/safety-sliders", handlers.GetAllSafetySlidersAdmin)
			admin.POST("/safety-sliders", handlers.CreateSafetySlider)
			admin.PUT("/safety-sliders/:id", handlers.UpdateSafetySlider)
			admin.DELETE("/safety-sliders/:id", handlers.DeleteSafetySlider)

			// ─── Supply Chain Admin ─────────────────────
			admin.GET("/supply-chain-strategies", handlers.GetAllSupplyChainStrategiesAdmin)
			admin.POST("/supply-chain-strategies", handlers.CreateSupplyChainStrategy)
			admin.PUT("/supply-chain-strategies/:id", handlers.UpdateSupplyChainStrategy)
			admin.DELETE("/supply-chain-strategies/:id", handlers.DeleteSupplyChainStrategy)

			admin.GET("/supply-chain-sustainability-items", handlers.GetAllSupplyChainSustainabilityItemsAdmin)
			admin.POST("/supply-chain-sustainability-items", handlers.CreateSupplyChainSustainabilityItem)
			admin.PUT("/supply-chain-sustainability-items/:id", handlers.UpdateSupplyChainSustainabilityItem)
			admin.DELETE("/supply-chain-sustainability-items/:id", handlers.DeleteSupplyChainSustainabilityItem)

			admin.GET("/supply-chain-policies", handlers.GetAllSupplyChainPoliciesAdmin)
			admin.POST("/supply-chain-policies", handlers.CreateSupplyChainPolicy)
			admin.PUT("/supply-chain-policies/:id", handlers.UpdateSupplyChainPolicy)
			admin.DELETE("/supply-chain-policies/:id", handlers.DeleteSupplyChainPolicy)

			// ─── CSR Vision Admin ───────────────────────
			admin.GET("/csr-vision-content", handlers.GetCsrVisionContent)
			admin.PUT("/csr-vision-content", handlers.UpdateCsrVisionContent)
			admin.GET("/csr-vision-strategies", handlers.GetAllCsrVisionStrategiesAdmin)
			admin.POST("/csr-vision-strategies", handlers.CreateCsrVisionStrategy)
			admin.PUT("/csr-vision-strategies/:id", handlers.UpdateCsrVisionStrategy)
			admin.DELETE("/csr-vision-strategies/:id", handlers.DeleteCsrVisionStrategy)

			// ─── CSR Report Admin ───────────────────────
			admin.GET("/csr-reports", handlers.GetAllCsrReportsAdmin)
			admin.POST("/csr-reports", handlers.CreateCsrReport)
			admin.PUT("/csr-reports/:id", handlers.UpdateCsrReport)
			admin.DELETE("/csr-reports/:id", handlers.DeleteCsrReport)

			admin.GET("/contact-info", handlers.GetContactInfo)
			admin.PUT("/contact-info", handlers.UpdateContactInfo)
			admin.GET("/contact-offices", handlers.GetAllContactOfficesAdmin)
			admin.POST("/contact-offices", handlers.CreateContactOffice)
			admin.PUT("/contact-offices/:id", handlers.UpdateContactOffice)
			admin.DELETE("/contact-offices/:id", handlers.DeleteContactOffice)

			admin.GET("/menus", handlers.GetAllMenus)
			admin.POST("/menus", handlers.CreateMenu)
			admin.PUT("/menus/:id", handlers.UpdateMenu)
			admin.DELETE("/menus/:id", handlers.DeleteMenu)
		}
	}

	return r
}
