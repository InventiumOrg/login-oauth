package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

var client *mongo.Client

func NewUserModel(mongo *mongo.Client) Models {
	client = mongo
	return Models{
		User: User{},
	}
}

type Models struct {
	User User
}

type User struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string    `bson:"username" json:"username"`
	Password  string    `bson:"password" json:"password"`
	Role      string    `bson:"role" json:"role"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func CreateNewUser(user User) error {
	collection := client.Database("users").Collection("users")

	_, err := collection.InsertOne(context.TODO(), User{
		Username:  user.Username,
		Password:  user.Password,
		Role:      user.Role,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		fmt.Println("Error creating new user: ", user.ID)
		return err
	}

	return nil

}
