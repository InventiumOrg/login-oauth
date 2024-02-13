package controllers

import (
	"login-oauth/models"
	"login-oauth/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.AuthService
	userService services.UserService
}

func NewAuthController(authService services.AuthService, userService services.UserService) AuthController {
	return AuthController{authService, userService}
}

func (ac *AuthController) SignUpUser(ctx *gin.Context) {
	var user *models.SignUpInput

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	if user.Password != user.ConfirmPassword {
		ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "passwords do not match"})
		return
	}

	newUser, err := ac.authService.SignUpUser(user)
	if err != nil {
		if strings.Contains(err.Error(), "username already exists") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "error", "message": "username already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": newUser})
}

// func (ac *AuthController) SignInUser(w http.ResponseWriter, r * http.Request) {
// 	var user *models.SignInInput
// }
