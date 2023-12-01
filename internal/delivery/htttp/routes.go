package htttp

import (
	"github.com/MikeMwita/messaging_app.git/internal/ports"
	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, repo ports.Repository) {
	messageHandler := NewMessageHandler(repo)

	r.GET("/messages", messageHandler.GetMessagesHandler)
	r.GET("/messages/:id", messageHandler.GetMessageHandler)
	r.POST("/messages", messageHandler.SendMessageHandler)
	r.POST("/simulate/send-message", messageHandler.SimulateSendMessageHandler)
	r.GET("/simulate/receive-messages", messageHandler.SimulateReceiveMessagesHandler)
}
