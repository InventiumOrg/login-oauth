package handlers

import (
	"login-oauth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestPayLoad struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"`
}

type UserPayLoad struct {
	Message string `json:"message"`
}

func SignUp(context *gin.Context) {
	var requestPayLoad RequestPayLoad

	if err := context.ShouldBindJSON(&requestPayLoad); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.CreateNewUser(models.User{
		Username: requestPayLoad.Username,
		Password: requestPayLoad.Password,
		Role:     requestPayLoad.Role,
	})

	if err != nil {
		context.JSON(http.StatusInternalServerError, UserPayLoad{
			Message: "Error signing up user",
		})
		return
	}

	userPayLoad := UserPayLoad{
		Message: "User has been signed up successfully",
	}

	context.JSON(http.StatusCreated, userPayLoad)
}
