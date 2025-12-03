package repository

import "gochat/internal/models"

type UserStore interface {
	GetUserByID(userID int) (*models.User, error)
	CreateUser(user *models.User) error
}
