package main

import (
	httpadapter "api/internal/adapters/http"
	"api/internal/adapters/persistence"
	"api/internal/application"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	_ "api/docs" // Import the generated Swagger docs

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Clube do Churrasco API
// @version 1.0
// @description API for managing users and events in the Clube do Churrasco application.
func main() {

	//it must be moved for some IoC container or initialization function|object
	// Load database configuration
	fmt.Println(os.Getwd())
	configFile, err := filepath.Abs("internal/adapters/persistence/config/db.yml")
	if err != nil {
		fmt.Println(os.Getwd())
		//log.Fatal(err)
	}
	dbConfig, err := persistence.NewDBConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new MongoDB client
	client, err := persistence.DB(*dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	// Get a handle to the database
	db := client.Database(dbConfig.Mongodb.Database)	

	// Repositories
	userRepo := persistence.NewMongoUserRepository(db, "users")
	eventRepo := persistence.NewMongoEventRepository(db, "events") // implement as needed

	// Services
	userService := application.NewUserService(userRepo)
	eventService := application.NewEventService(eventRepo) // implement as needed

	// Handlers
	handler := httpadapter.NewHandler(eventService, userService) // pass eventService when implemented

	// Gin setup
	r := gin.Default()
	//docs.SwaggerInfo.BasePath = "/api/v1" // Set the base path for Swagger
	//v1 := r.Group("/api/v1")
	//{
		// Register Swagger routes
	//	v1.POST("/users", handler.CreateUser) // Example route, replace with actual routes

	//}
	//r.GET("/swagger/v1/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler(), ginSwagger.InstanceName("v1")))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	httpadapter.RegisterRoutes(r, handler)
	r.Run()
}
