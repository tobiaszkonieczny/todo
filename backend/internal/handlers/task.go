package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
	"github.com/tobiaszkonieczny/todo.git/internal/ws"
)

// GetTasks godoc
// @Summary      List of all tasks
// @Description  Retrieve a list of all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}   models.Task
// @Router       /tasks [get]
func GetTasks(c *gin.Context) {
	var tasks []models.Task

	// Preload attachments
	if err := models.DB.Preload("Attachments").Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

// CreateTask godoc
// @Summary      Create a new task
// @Description  Add a new task to the list
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task  body      models.Task  true  "Task to create"
// @Success      201   {object}  models.Task
// @Failure      400   {object}  map[string]string
// @Router       /tasks [post]
func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//userId should be in gin context
	val, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userID := uint(val.(float64))
	task.UserID = userID

	if task.CategoryID != nil {
		var category models.Category
		if err := models.DB.First(&category, *task.CategoryID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid category_id"})
			return
		}
	}

	if err := models.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := fmt.Sprintf(`{"event":"new_task","task_id":%d,"title":"%s"}`, task.ID, task.Title)
	ws.Broadcast([]byte(msg))

	c.JSON(http.StatusCreated, task)
}

// UpdateTask godoc
// @Summary      Update a task
// @Description  Update an existing task by ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id    path      int          true  "Task ID"
// @Param        task  body      models.Task  true  "Updated task"
// @Success      200   {object}  models.Task
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /tasks/{id} [put]
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
		return
	}
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	models.DB.Save(&task)
	c.JSON(http.StatusOK, task)
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Delete a task by ID
// @Tags         tasks
// @Produce      json
// @Param        id    path      int  true  "Task ID"
// @Success      204   "No Content"
// @Failure      404   {object}  map[string]string
// @Router       /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	models.DB.Delete(&models.Task{}, id)
	c.Status(http.StatusNoContent)
}
