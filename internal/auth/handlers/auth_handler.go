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

func (h *AuthHandler) RegisterAdmin(ctx *gin.Context) {

	var request dto.RegisterRequest

	//should bind this like using req.body like use node
	//here we dont use ShouldBindJSON it will pass the empty dto.RegisterRequest and db operarion will fail

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.AuthResponse{Message: "Inavlid body Request"})
	}
	//Gin context has AbortWithStatusJSON it will stop the whole function
	//like simple try catch where we cathed error and retrn
	//for more
	//https://pkg.go.dev/github.com/gin-gonic/gin#Context
	if err := h.service.Register(request.Username, request.Email, request.Password, "admin"); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.AuthResponse{Message: err.Error()})
		return
	}
	//this net http module will have an abstracitn of http code here
	//it will send 201 code there are mapping https://go.dev/src/net/http/status.go pls have a look
	ctx.JSON(http.StatusCreated, dto.AuthResponse{Message: "Register Success"})

}

func (h *AuthHandler) Login(ctx *gin.Context)  {
	/**
	*see this is a serailition & deserailaton
	*we often see the clinet will interact with the JSON to api acuse browser env will have js% json
	*but out Golang will only impliclty untestnad structs this validation is verty imprtant
	 */
	var request dto.LoginRequest
	if err := ctx.ShouldBindBodyWithJSON(&request); err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, dto.AuthResponse{Message: "Inavlid body Request"})

	}

	token, err := h.service.Login(request)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, dto.AuthResponse{Message: err.Error()})

	}

	//? res.setCookie()
	ctx.SetCookie("Authorization", token, 3600, "/", "localhost", false, true)
	
	//? res.status(200).json(result)
	ctx.JSON(http.StatusOK,dto.LoginResponse{Token: token})
}
