package ports

import "github.com/MikeMwita/messaging_app.git/internal/models"

type Repository interface {
	GetMessages() ([]models.Message, error)
	GetMessage(id int) (models.Message, error)
	GetUser(name string) (models.User, error)
	InsertMessage(message models.Message) error
	SeedMessagesFromCSV(filepath string) error
}
