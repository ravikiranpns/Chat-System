package handlers

import (
	"chat-system/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func SendMessage(c *gin.Context) {
	var message struct {
		Sender    string `json:"sender"`
		Recipient string `json:"recipient"`
		Content   string `json:"content"`
	}
	if err := c.BindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	id := uuid.New()
	timestamp := time.Now()

	err := db.GetSession().Query(`INSERT INTO chat.messages (id, sender, recipient, content, timestamp) VALUES (?, ?, ?, ?, ?)`,
		id, message.Sender, message.Recipient, message.Content, timestamp).Exec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "message sent"})
}

func GetMessages(c *gin.Context) {
	username := c.Query("username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return
	}

	var messages []struct {
		ID        uuid.UUID `json:"id"`
		Sender    string    `json:"sender"`
		Recipient string    `json:"recipient"`
		Content   string    `json:"content"`
		Timestamp time.Time `json:"timestamp"`
	}

	iter := db.GetSession().Query(`SELECT id, sender, recipient, content, timestamp FROM chat.messages WHERE recipient = ?`, username).Iter()
	for iter.Scan(&messages) {
		messages = append(messages, messages)
	}

	if err := iter.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve messages"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"messages": messages})
}
