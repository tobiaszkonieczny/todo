package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

// GetCategories godoc
// @Summary Get all categories
// @Description Retrieve a list of all available categories for organizing tasks
// @Tags categories
// @Produce json
// @Security BearerAuth
// @Success 200 {array} models.Category "List of categories retrieved successfully"
// @Failure 500 {object} map[string]string "Internal server error" example({"error": "database error"})
// @Router /categories [get]
func GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := models.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// CreateCategory godoc
// @Summary Create new category
// @Description Create a new category for organizing tasks. Category name must be unique.
// @Tags categories
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param category body models.Category true "Category data (only name field is required)"
// @Success 201 {object} models.Category "Category created successfully"
// @Failure 400 {object} map[string]string "Bad request - invalid input" example({"error": "invalid JSON format"})
// @Failure 500 {object} map[string]string "Internal server error - possibly duplicate name" example({"error": "category name already exists"})
// @Router /categories [post]
func CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := models.DB.Create(&category).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, category)
}
