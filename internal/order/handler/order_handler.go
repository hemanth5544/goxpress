package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/internal/auth/model"
	"github.com/hemanth5544/goxpress/internal/order/dto"
	"github.com/hemanth5544/goxpress/internal/order/services"
)

type OrderHandler struct {
	orderService *services.OrderServices
}

func NewOrderHandler(orderService *services.OrderServices) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func (h *OrderHandler) Checkout(c *gin.Context) {

	user, exist := c.Get("user")

	if !exist {
		c.JSON(http.StatusUnauthorized, dto.OrderMessage{
			Message: "user is unauthorized",
		})
	}
	//?Bindin ht user touse model so that we can map id or any thing tath mapped in user model

	userID, ok := user.(model.User)
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.OrderMessage{
			Message: "Invalid user",
		})
	}

	var request dto.CheckoutRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, dto.OrderMessage{
			Message: "Invalid payment method",
		})
	}

	err := h.orderService.CheckoutService(userID.ID, request.PaymentMethod)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.OrderMessage{
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.OrderMessage{
		Message: "Checkout successful",
	})

}
