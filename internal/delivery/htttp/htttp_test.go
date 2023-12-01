package htttp

import (
	"encoding/json"
	"github.com/MikeMwita/messaging_app.git/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockRepository struct {
}

func (r *MockRepository) GetMessages() ([]models.Message, error) {
	return []models.Message{{ID: 1, UserID: 100, Content: "Test message"}}, nil
}

func (r *MockRepository) GetMessage(id int) (models.Message, error) {

	return models.Message{ID: 1, UserID: 100, Content: "Test message"}, nil
}

func (r *MockRepository) InsertMessage(message models.Message) error {

	return nil
}

func TestGetMessagesHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	//handler := http.NewMessageHandler(&MockRepository{})

	req, err := http.NewRequest(http.MethodGet, "/messages", nil)
	assert.NoError(t, err)
	rr := httptest.NewRecorder()
	router := gin.Default()
	//router.GET("/messages", handler.GetMessagesHandler)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var messages []models.Message
	err = json.NewDecoder(rr.Body).Decode(&messages)
	assert.NoError(t, err)

	// Check the content of the response
	assert.Len(t, messages, 1)
	assert.Equal(t, 1, messages[0].ID)
	assert.Equal(t, 100, messages[0].UserID)
	assert.Equal(t, "Test message", messages[0].Content)
}
