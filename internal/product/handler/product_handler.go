package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/internal/product/dto"
	"github.com/hemanth5544/goxpress/internal/product/model"
	"github.com/hemanth5544/goxpress/internal/product/services"
)

type ProductHandler struct {
	services *services.ProductServices
}

func NewProductHandler(services *services.ProductServices) *ProductHandler {
	return &ProductHandler{services: services}
}

//? in Handler Layer we play more wiht gin
//* Coz the req and res will be directly having form route to Handler only

func (h *ProductHandler) CreateProduct(c *gin.Context) {

	var productRequest dto.ProductRequest

	if err := c.ShouldBindBodyWithJSON(&productRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Invalid body request",
		})
		return
	}

	err := h.services.CreateProduct(productRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "failed to create product",
		})
		return
	}
	//? see when we retun wiht res.json this gin context wants a response body also
	/*
		//*if we need to send only status wiht out any drama of Json use c.Status similar to res.status()
		//* c.Status(http.StatusOK)
	*/

	c.JSON(http.StatusOK, dto.Message{Message: "Created New Product"})

}

func (h *ProductHandler) DeleteProductById(c *gin.Context) {
	//? we need to contever to sting to int
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Failed to get id",
		})

	}

	if err := h.services.DeleteProductById(id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Failed to delete product",
		})
	}

	c.JSON(http.StatusOK, dto.Message{
		Message: "Success delete product",
	})
}

func (h *ProductHandler) GetProductById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Failed to get id",
		})

	}

	product, err := h.services.GetProductById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Failed to get product",
		})
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		Message: "Success getting product",
		Data:    *product,
	})
}

func (h *ProductHandler) GetAllProduct(c *gin.Context) {
	products, err := h.services.GetAllProduct()

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Failed to get all product",
		})
	}

	c.JSON(http.StatusOK, dto.ProductResponseList{
		Message: "Successfully getting all product",
		Data:    *products,
	})
}

func (h *ProductHandler) UpdateProductById(c *gin.Context) {
	productID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Message{
			Message: "Invalid product id",
		})
		return
	}

	var updateRequest model.Product

	//Binding with Json
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{
			Message: "Invalid Request body",
		})
	}

	updatedProduct, err := h.services.UpdateProductById(uint(productID), updateRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Message{Message: err.Error()})
	}

	c.JSON(http.StatusOK, dto.ProductResponse{
		Message: "Successfully update product",
		Data:    *updatedProduct,
	})

}
