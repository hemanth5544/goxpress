package cart

import (
	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/internal/cart/handler"
	"github.com/hemanth5544/goxpress/internal/cart/repository"
	"github.com/hemanth5544/goxpress/internal/cart/services"
	"github.com/hemanth5544/goxpress/middleware"
	"gorm.io/gorm"

	productRepo "github.com/hemanth5544/goxpress/internal/product/repository"
)

func SetupCart(router *gin.Engine, db *gorm.DB) {

	//? by using this services and passing hte depnceids here can reduce many cahnges in serives
	//! beetter and easy to wirte testcase and buiild app scaliable

	//*	DB → Repository → Service → Handler → Routes

	repo := repository.NewCartRepository(db)
	productRepo := productRepo.NewProductRepository(db)
	services := services.NewCartServices(repo, productRepo)
	cartHandler := handler.NewCartHandler(services)
	//just grouping wiht version and cart prefix
	//! this can prevent the version mixups
	cart := router.Group("api/v1/cart")
	{
		cart.GET("/", middleware.RoleMiddleware("user"), cartHandler.GetCart)
		cart.POST("/add", middleware.RoleMiddleware("user"), cartHandler.AddToCart)
		cart.DELETE("/item/:id", middleware.RoleMiddleware("user"), cartHandler.RemoveItem)
		cart.PUT("/item/:id", middleware.RoleMiddleware("user"), cartHandler.UpdateQuantity)
	}

}
