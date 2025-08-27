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
// @description API for tasks management.
// @host localhost:8081
// @BasePath /
func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Println("⚠️ Missing env file, proceeding with system environment variables")
	}

	// Invoke database connection
	models.ConnectDatabase()

	// Router Gin
	r := gin.Default()

	// Endpoints CRUD
	api := r.Group("/tasks")
	api.Use(middleware.JWTAuth())
	{
		api.GET("/", handlers.GetTasks)
		api.POST("/new", handlers.CreateTask)
		api.PUT("update/:id", handlers.UpdateTask)
		api.DELETE("delete/:id", handlers.DeleteTask)
	}
	category := r.Group("/categories")
	category.Use(middleware.JWTAuth())
	{
		category.GET("/", handlers.GetCategories)
		category.POST("/new", handlers.CreateCategory)
	}

	// Endpoints auth
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/logout", handlers.Logout)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start serwera
	r.Run(":8081")
}
