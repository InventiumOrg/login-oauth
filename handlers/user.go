package handlers

import (
	"fmt"
	"login-oauth/config"
	"login-oauth/models"
	"login-oauth/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type RequestPayLoad struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role,omitempty" binding:"required,omitempty"`
}

type UserPayLoad struct {
	Message string `json:"message"`
}

type PayLoad struct {
	Sub string
	Iss string
	AUD string
}

func SignUp(context *gin.Context) {
	var requestPayLoad RequestPayLoad

	if err := context.ShouldBindJSON(&requestPayLoad); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	encryptedPassword, _ := utils.HashPassword(requestPayLoad.Password)

	err := models.CreateNewUser(models.User{
		Username: requestPayLoad.Username,
		Password: encryptedPassword,
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

func Signin(context *gin.Context) {
	var requestPayLoad RequestPayLoad

	if err := context.ShouldBindJSON(&requestPayLoad); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// encryptedPassword, _ := utils.HashPassword(requestPayLoad.Password)

	matchedUser := models.GetUserByUsername(requestPayLoad.Username)
	err := utils.VerifyPassword(matchedUser.Password, requestPayLoad.Password)
	if err != nil {
		context.JSON(http.StatusUnauthorized, UserPayLoad{
			Message: "Invalid Crendeitials",
		})
	} else {
		payLoad := PayLoad{
			Sub: matchedUser.ID,
			Iss: "inventium",
			AUD: matchedUser.Role,
		}
		config, err := config.LoadConfig(".")
		token, err := utils.CreateToken(60*time.Minute, payLoad, config.PrivateKey)
		fmt.Println(token)
		if err != nil {
			fmt.Println(err)
		}
		context.SetCookie("token", token, 3600, "/", "localhost", false, true)
		context.JSON(http.StatusOK, UserPayLoad{
			Message: "Login Success",
		})
		return
	}

}
