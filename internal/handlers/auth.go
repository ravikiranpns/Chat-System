package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"chat-system/internal/db"
)

func Register(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	session := db.GetSession()
	if session == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no Cassandra session"})
		return
	}

	err := session.Query("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password).Exec()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered"})
}

func Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	session := db.GetSession()
	if session == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no Cassandra session"})
		return
	}

	var password string
	err := session.Query("SELECT password FROM users WHERE username = ?", user.Username).Scan(&password)
	if err != nil || password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user logged in"})
}
