package routes

import (
	"login-oauth/controllers"
	"login-oauth/middlewares"
	"login-oauth/services"

	"github.com/gin-gonic/gin"
)

type UserRouteController struct {
	userController controllers.UserController
}

func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup, userService services.UserService) {

	router := rg.Group("users")
	router.Use(middlewares.DeserializeUser(userService))
	router.GET("/me", uc.userController.GetMe)
}
