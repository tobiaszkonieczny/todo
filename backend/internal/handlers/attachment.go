package handlers

import (
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

// UploadAttachment godoc
// @Summary Upload file attachment to a task
// @Description Upload a file attachment to a specific task. The file will be stored in the database and associated with the task.
// @Tags attachments
// @Accept multipart/form-data
// @Produce json
// @Param task_id path int true "Task ID to attach file to"
// @Param file formData file true "File to upload"
// @Security BearerAuth
// @Success 201 {object} map[string]interface{} "File uploaded successfully" example({"message": "file uploaded", "attachment_id": 123})
// @Failure 400 {object} map[string]string "Bad request - no file uploaded" example({"error": "no file uploaded"})
// @Failure 404 {object} map[string]string "Task not found" example({"error": "task not found"})
// @Failure 500 {object} map[string]string "Internal server error" example({"error": "database error"})
// @Router /tasks/{task_id}/attachments [post]
func UploadAttachment(c *gin.Context) {
	taskID := c.Param("task_id")

	var task models.Task
	if err := models.DB.First(&task, taskID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no file uploaded"})
		return
	}

	f, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer func(f multipart.File) { // ensure file is closed after reading (wrapping func returns)
		err := f.Close()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	}(f)

	content := make([]byte, file.Size)
	_, err = f.Read(content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	attachment := models.Attachment{
		TaskID:      task.ID,
		FileName:    file.Filename,
		Content:     content,
		ContentType: file.Header.Get("Content-Type"),
		Size:        file.Size,
	}

	if err := models.DB.Create(&attachment).Error; err != nil { // save attachment to DB and simultaneously handle errors
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "file uploaded", "attachment_id": attachment.ID})
}

// DownloadAttachment godoc
// @Summary Download file attachment
// @Description Download a file attachment by its ID. Returns the raw file content with appropriate content type.
// @Tags attachments
// @Produce application/octet-stream
// @Produce image/jpeg
// @Produce image/png
// @Produce application/pdf
// @Produce text/plain
// @Param attachment_id path int true "Attachment ID to download"
// @Security BearerAuth
// @Success 200 {file} file "File content"
// @Failure 404 {object} map[string]string "Attachment not found" example({"error": "attachment not found"})
// @Router /attachments/{attachment_id}/download [get]
func DownloadAttachment(c *gin.Context) {
	attachmentID := c.Param("attachment_id")

	var att models.Attachment
	if err := models.DB.First(&att, attachmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "attachment not found"})
		return
	}

	c.Data(http.StatusOK, att.ContentType, att.Content)
}
