package route

import (
	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/controller/apikey"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/controller/auth"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/controller/user"
	"github.com/oadultradeepfield/galaxy10-apikey/backend/internal/middleware"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	authController := auth.NewAuthController(db)
	apiKeyController := apikey.NewAPIKeyController(db)
	userController := user.NewUserController(db)

	r.Use(middleware.CorsMiddleware())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	public := r.Group("/api")
	{
		auth := public.Group("/auth")
		{
			auth.GET("/google/signin", authController.GoogleSignin)
			auth.GET("/google/callback", authController.GoogleCallback)
		}
	}

	protected := r.Group("/api")
	protected.Use(middleware.JWTMiddleware(db))
	{
		apikey := protected.Group("/apikeys")
		{
			apikey.GET("/", apiKeyController.GetAllAPIKeys)
			apikey.POST("/", apiKeyController.CreateAPIKey)
			apikey.PUT("/:id", apiKeyController.UpdateAPIKey)
			apikey.DELETE("/:id", apiKeyController.DeleteAPIKey)
		}

		user := protected.Group("/user")
		{
			user.GET("/", userController.GetCurrentUserInfo)
		}
	}
}
