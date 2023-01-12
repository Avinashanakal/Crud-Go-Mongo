package services

import "github.com/Avinashanakal/models"

// implement methods using struct for below mentioned api contracts
type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
