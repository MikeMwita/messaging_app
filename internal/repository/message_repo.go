package repository

import (
	"database/sql"
	"encoding/csv"
	"github.com/MikeMwita/messaging_app.git/internal/models"
	"github.com/MikeMwita/messaging_app.git/internal/ports"
	"io"
	"os"
	"strconv"
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetMessages() ([]models.Message, error) {
	rows, err := r.db.Query("select * from messages order by time")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var messages []models.Message
	for rows.Next() {
		var message models.Message
		err = rows.Scan(&message.ID, &message.UserID, &message.Time, &message.Content)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}
	return messages, nil
}

func (r *repository) GetMessage(id int) (models.Message, error) {
	row := r.db.QueryRow("select * from messages where id = ?", id)
	var message models.Message
	err := row.Scan(&message.ID, &message.UserID, &message.Time, &message.Content)
	if err != nil {
		return message, err
	}
	return message, nil
}

func (r *repository) GetUser(name string) (models.User, error) {
	row := r.db.QueryRow("select * from users where name = ?", name)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Role, &user.Image)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) InsertMessage(message models.Message) error {
	_, err := r.db.Exec("insert into messages (user_id, time, content) values (?, ?, ?)", message.UserID, message.Time, message.Content)
	if err != nil {
		return err
	}
	return nil
}

// SeedMessagesFromCSV reads messages from a CSV file and inserts them into the database
func (r *repository) SeedMessagesFromCSV(filepath string) error {
	messages, err := readMessagesFromCSV(filepath)
	if err != nil {
		return err
	}

	for _, message := range messages {
		err := r.InsertMessage(message)
		if err != nil {
			return err
		}
	}

	return nil
}

// readMessagesFromCSV is a helper function to read messages from a CSV file
func readMessagesFromCSV(filepath string) ([]models.Message, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	var messages []models.Message

	// Loop over the records and parse them as Message structs
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		userID, err := strconv.Atoi(record[0])
		if err != nil {
			return nil, err
		}

		message := models.Message{
			ID:      userID,
			UserID:  userID,
			Time:    record[1],
			Content: record[2],
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func NewRepository(db *sql.DB) ports.Repository {
	return &repository{db: db}
}
