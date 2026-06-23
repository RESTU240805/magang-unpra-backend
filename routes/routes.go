package routes

import (
	"time"

	"magang-unpra-backend/handlers"
	"magang-unpra-backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "X-Requested-With"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Serve uploaded files
	r.Static("/uploads", "./uploads")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/api/login", handlers.Login)
	r.POST("/api/upload", handlers.UploadImage)

	api := r.Group("/api")
	{
		// Public
		api.GET("/news", handlers.GetAllNews)
		api.GET("/news/:id", handlers.GetNewsById)
		api.GET("/products", handlers.GetAllProducts)
		api.GET("/products/:id", handlers.GetProductById)
		api.GET("/company-profile", handlers.GetCompanyProfile)
		api.GET("/about-section", handlers.GetAboutSection)
		api.PUT("/about-section", handlers.UpdateAboutSection)
		api.GET("/product-slides", handlers.GetAllSlides)
		api.GET("/product-page", handlers.GetProductPage)
		api.PUT("/product-page", handlers.UpdateProductPage)
		api.GET("/team-members", handlers.GetAllTeamMembers)
		api.GET("/community-cards", handlers.GetAllCommunityCards)
		api.GET("/creeds", handlers.GetAllCreeds)
		api.GET("/company-documents", handlers.GetAllCompanyDocuments)

		// Admin
		admin := api.Group("/admin")
		admin.Use(middleware.AuthRequired())
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

			api.GET("/org-structure", handlers.GetOrgStructure)
			api.GET("/org-chart", handlers.GetOrgChart)

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

			api.GET("/menus", handlers.GetActiveMenus)

			admin.GET("/menus", handlers.GetAllMenus)
			admin.POST("/menus", handlers.CreateMenu)
			admin.PUT("/menus/:id", handlers.UpdateMenu)
			admin.DELETE("/menus/:id", handlers.DeleteMenu)
		}
	}

	return r
}
