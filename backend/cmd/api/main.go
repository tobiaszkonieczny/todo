package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tobiaszkonieczny/todo.git/internal/middleware"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tobiaszkonieczny/todo.git/docs"

	"github.com/tobiaszkonieczny/todo.git/internal/handlers"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

// @title ToDo API
// @version 1.0
// @description API do zarządzania zadaniami (CRUD).
// @host localhost:8081
// @BasePath /
func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Println("⚠️  Brak pliku .env, używam zmiennych środowiskowych")
	}

	// Połączenie z bazą
	models.ConnectDatabase()

	// Router Gin
	r := gin.Default()

	// Endpointy CRUD
	api := r.Group("/tasks")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/", handlers.GetTasks)
		api.POST("/", handlers.CreateTask)
		api.PUT("/:id", handlers.UpdateTask)
		api.DELETE("/:id", handlers.DeleteTask)
	}

	// Endpointy auth
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start serwera
	r.Run(":8081")
}
