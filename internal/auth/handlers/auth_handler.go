package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hemanth5544/goxpress/internal/auth/dto"
	"github.com/hemanth5544/goxpress/internal/auth/services"
)

type AuthHandler struct {
	service *services.AuthService
}

func smaple(dto.AuthResponse) {

}

func (h *AuthHandler) RegisterUser(ctx *gin.Context) {

	var request dto.RegisterRequest
	//validation in node we user try call
	//but in Go we mainly diffnetn on if and err !=nil
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.AuthResponse{Message: "Inavlid body Request"})
	}

	if err := h.service.Register(request.Username, request.Email, request.Password, "user"); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.AuthResponse{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, dto.AuthResponse{Message: "Register Success"})
}
