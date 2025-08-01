package main

import (
	httpadapter "api/internal/adapters/http"
	"api/internal/adapters/persistence"
	"api/internal/application"
	"context"
	"log"	
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB connection
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("clube_do_churrasco")

	// Repositories
	userRepo := persistence.NewMongoUserRepository(db, "users")
	// eventRepo := persistence.NewMongoEventRepository(db, "events") // implement as needed

	// Services
	userService := application.NewUserService(userRepo)
	// eventService := application.NewEventService(eventRepo) // implement as needed

	// Handlers
	handler := httpadapter.NewHandler(nil, userService) // pass eventService when implemented

	// Gin setup
	r := gin.Default()
	httpadapter.RegisterRoutes(r, handler)
	r.Run()
}

// Event represents the event model
type Event struct {
	ID          string    `json:"id"`
	Date        time.Time `json:"date" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Attenders   []string  `json:"attenders" binding:"required"`
	Details     string    `json:"details"`
}

// User represents the user model
type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	WhatsApp string `json:"whatsapp" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// In-memory event storage (for demonstration)
var events = []Event{}

// In-memory user storage (for demonstration)
var users = []User{}



