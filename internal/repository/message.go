package repository

import "gochat/internal/models"

type MessageStore interface {
	GetMessagesByUserID(userID int) ([]*models.Message, error)
	CreateMessage(message *models.Message) error
}
