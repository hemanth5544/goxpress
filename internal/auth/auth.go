package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/internal/auth/handlers"
	"github.com/hemanth5544/goxpress/internal/auth/repository"
	"github.com/hemanth5544/goxpress/internal/auth/services"
	"github.com/hemanth5544/goxpress/middleware"

	"gorm.io/gorm"
)

func SetupAuth(router *gin.Engine, db *gorm.DB) {
	authRepo := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	auth := router.Group("api/v1/auth")

	{
		auth.POST("/login", authHandler.Login)
		admin := router.Group("/admin")
		{
			admin.POST("/register", authHandler.RegisterAdmin)
			admin.GET("/dashboard", middleware.RoleMiddleware("admin"), func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"message": "this is admin"})
			})
		}

		user := router.Group("/user")
		{
			user.POST("/register", authHandler.RegisterUser)
		}
	}
}
