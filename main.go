package main

import (
	"context"
	"fmt"
	"log"
	"login-oauth/config"
	"login-oauth/models"
	"login-oauth/routes"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type App struct {
	Router *gin.Engine
}

var (
	client *mongo.Client
	// redisclient    *redis.Client
	// authCollection *mongo.Collection
)

func main() {
	// Load Config
	config, err := config.LoadConfig(".")

	if err != nil {
		fmt.Println("Error loading config")
	}

	// connect to mongodb to perform CRUD on User
	mongoClient, err := connectToMongo(config)

	if err != nil {
		log.Panic(err)
	}
	client = mongoClient
	models.NewUserModel(client)
	// connect to redis to save session

	app := App{
		Router: routes.Routes(),
	}

	app.Router.Run(config.Port)
}

func connectToMongo(config config.Config) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	clientOption := options.Client().ApplyURI(config.DBUri)
	clientOption.SetAuth(options.Credential{
		Username: config.DBUsername,
		Password: config.DBPassword,
	})
	client, err := mongo.Connect(ctx, clientOption)
	if err != nil {
		fmt.Println("Error connecting to mongo")
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return client, nil
}
