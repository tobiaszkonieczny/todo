package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

// GetTasks godoc
// @Summary      Get all tasks
// @Description  Retrieve a list of all tasks
// @Tags         tasks
// @Produce      json
// @Success      200  {array}   models.Task
// @Router       /tasks [get]
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	models.DB.Find(&tasks)
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
	models.DB.Create(&task)
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
