package htttp

import (
	"database/sql"
	"github.com/MikeMwita/messaging_app.git/internal/models"
	"github.com/MikeMwita/messaging_app.git/internal/ports"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MessageHandler struct {
	Repo ports.Repository
}

type Repository struct {
	Db *sql.DB
}

// GetMessagesHandler handles GET requests to retrieve messages
func (h *MessageHandler) GetMessagesHandler(c *gin.Context) {
	messages, err := h.Repo.GetMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

// GetMessageHandler handles GET requests to retrieve a specific message by ID
func (h *MessageHandler) GetMessageHandler(c *gin.Context) {
	id := c.Param("id")

	messageID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	message, err := h.Repo.GetMessage(messageID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Message not found"})
		return
	}

	c.JSON(http.StatusOK, message)
}

// SendMessageHandler handles POST requests to send a message
func (h *MessageHandler) SendMessageHandler(c *gin.Context) {
	var newMessage models.Message
	if err := c.ShouldBindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.Repo.InsertMessage(newMessage)
	if err != nil {
		// Handle the error and respond with an appropriate status code
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert message"})
		return
	}

	c.JSON(http.StatusOK, newMessage)
}

// SimulateSendMessageHandler simulates sending a message

func (h *MessageHandler) SimulateSendMessageHandler(c *gin.Context) {
	var newMessage models.Message
	if err := c.ShouldBindJSON(&newMessage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.Repo.InsertMessage(newMessage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert message"})
		return
	}

	c.JSON(http.StatusOK, newMessage)
}

func (h *MessageHandler) SimulateReceiveMessagesHandler(c *gin.Context) {
	messages := []models.Message{}

	query := "SELECT user_id, content FROM messages"

	rows, err := h.Repo.GetDB().Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query messages"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		message := models.Message{}

		err := rows.Scan(&message.UserID, &message.Content)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan message"})
			return
		}

		messages = append(messages, message)
	}

	err = rows.Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read messages"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func NewMessageHandler(repo ports.Repository) *MessageHandler {
	return &MessageHandler{Repo: repo}
}
