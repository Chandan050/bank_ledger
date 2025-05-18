package main

import (
	"context"
	"log"
	"os"

	"banking_ledger/internal/api/handler"
	"banking_ledger/internal/api/router"
	MongoDB "banking_ledger/internal/repository/mongodb"
	"banking_ledger/internal/service"
	"banking_ledger/queue"

	"banking_ledger/internal/repository/postgres"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	//database init
	postgres, err := postgres.NewPostgresAccountRepository()
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	mongoDB, err := MongoDB.NewMongoDBAccountRepository()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	//service init
	service := service.NewService(postgres, mongoDB)

	//consumer init
	go queue.ConsumeTransactions(context.TODO(), *service, "createTransaction", 0)

	// Initialize router
	r := gin.Default()

	// Start server
	h := handler.NewHandler(service)
	router.SetupRouter(r, h)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
