package order

import (
	"github.com/gin-gonic/gin"
	cartRepo "github.com/hemanth5544/goxpress/internal/cart/repository"
	"github.com/hemanth5544/goxpress/internal/order/handler"
	"github.com/hemanth5544/goxpress/internal/order/repository"
	"github.com/hemanth5544/goxpress/internal/order/services"
	productRepo "github.com/hemanth5544/goxpress/internal/product/repository"
	"github.com/hemanth5544/goxpress/middleware"
	"gorm.io/gorm"
)

func SetupOrder(router *gin.Engine, db *gorm.DB) {
	orderRepo := repository.NewOrderRepository(db)
	productRepo := productRepo.NewProductRepository(db)
	cartRepo := cartRepo.NewCartRepository(db)
	orderService := services.NewOrderServices(orderRepo, productRepo, cartRepo)
	orderHandler := handler.NewOrderHandler(orderService)

	router.POST("api/v1/order/checkout", middleware.RoleMiddleware("user"), orderHandler.Checkout)
}
