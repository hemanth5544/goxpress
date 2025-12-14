package product

import (
	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/internal/product/handler"
	"github.com/hemanth5544/goxpress/internal/product/repository"
	"github.com/hemanth5544/goxpress/internal/product/services"
	"github.com/hemanth5544/goxpress/middleware"
	"gorm.io/gorm"
)

func ProductRouter(router *gin.Engine, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productServices := services.NewProductServices(productRepo)
	productHandler := handler.NewProductHandler(productServices)
	product := router.Group("api/v1/product")

	{
		product.POST("/", middleware.RoleMiddleware("admin"), productHandler.CreateProduct)
		product.GET("/:id", productHandler.GetProductById)
		product.GET("/", productHandler.GetAllProduct)
		product.PUT("/:id", middleware.RoleMiddleware("admin"), productHandler.UpdateProductById)
		product.DELETE("/:id", middleware.RoleMiddleware("admin"), productHandler.DeleteProductById)
	}

}
