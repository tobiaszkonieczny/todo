package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() //execute the handler

		// after request
		log := models.Log{
			Method:    c.Request.Method,
			Path:      c.Request.URL.Path,
			IP:        c.ClientIP(),
			UserAgent: c.Request.UserAgent(),
			CreatedAt: start,
		}

		models.DB.Create(&log)
	}
}
