package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"chat-system/internal/db"
	"chat-system/internal/handlers"
)

func main() {
	cassandraHost := os.Getenv("CASSANDRA_HOST")
	cassandraPort := os.Getenv("CASSANDRA_PORT")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	if err := db.InitCassandra(cassandraHost, cassandraPort); err != nil {
		log.Fatalf("failed to connect to Cassandra: %v", err)
	}
	defer db.GetSession().Close()

	if err := db.InitRedis(redisHost, redisPort); err != nil {
		log.Fatalf("failed to connect to Redis: %v", err)
	}

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login)
	r.POST("/send", handlers.SendMessage)
	r.GET("/messages", handlers.GetMessages)

	r.Run(":8080")
}
