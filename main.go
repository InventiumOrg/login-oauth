package main

import (
	"fmt"
	"login-oauth/config"
	"login-oauth/routes"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
}

// var (
// 	server         *gin.Engine
// 	ctx            context.Context
// 	mongoclient    *mongo.Client
// 	redisclient    *redis.Client
// 	authCollection *mongo.Collection
// )

func main() {
	// Load Config
	config, err := config.LoadConfig(".")

	if err != nil {
		fmt.Println("Error loading config")
	}

	// connect to mongodb to perform CRUD on User
	// connect to redis to save session

	app := App{
		Router: routes.Routes(),
	}
	app.Router.Run(config.Port)
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{"http://localhost:8000", "http://localhost:3000"}
	// corsConfig.AllowCredentials = true

}
