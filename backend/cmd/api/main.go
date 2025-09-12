package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tobiaszkonieczny/todo.git/internal/middleware"
	"github.com/tobiaszkonieczny/todo.git/internal/ws"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/tobiaszkonieczny/todo.git/docs"
	_ "github.com/tobiaszkonieczny/todo.git/internal/handlers"

	"github.com/tobiaszkonieczny/todo.git/internal/handlers"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

// @title ToDo API
// @version 1.0
// @description API for tasks management.
// @host localhost:8081
// @BasePath /
func main() {

	// Invoke database connection
	models.ConnectDatabase()

	// Router Gin
	r := gin.Default()
	r.Use(middleware.RequestLogger())

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
	attachment := r.Group("/attachments")
	attachment.Use(middleware.JWTAuth())
	{
		attachment.POST("/upload/:task_id", handlers.UploadAttachment)
		attachment.GET("/download/:attachment_id", handlers.DownloadAttachment)
	}

	//Websocket
	r.GET("/ws", func(c *gin.Context) {
		ws.HandleWS(c.Writer, c.Request)
	})

	// Endpoints auth
	r.POST("/auth/register", handlers.Register)
	r.POST("/auth/login", handlers.Login)
	r.POST("/auth/logout", handlers.Logout)

	// Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start serwera
	r.Run(":8081")
}
