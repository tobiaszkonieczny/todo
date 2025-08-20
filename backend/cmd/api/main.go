package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tobiaszkonieczny/todo.git/internal/handlers"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

func main() {
	if err := godotenv.Load("configs/.env"); err != nil {
		log.Println("⚠️  Brak pliku .env, używam zmiennych środowiskowych")
	}

	models.ConnectDatabase()

	r := gin.Default()

	r.GET("/tasks", handlers.GetTasks)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	r.Run(":8080")
}
