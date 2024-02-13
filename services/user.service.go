package services

import "login-oauth/models"

type UserService interface {
	FindUserById(string) (*models.DBResponse, error)
	FindUserByUsername(string) (*models.DBResponse, error)
}
