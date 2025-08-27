package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tobiaszkonieczny/todo.git/internal/models"
)

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
	defer f.Close()

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

	if err := models.DB.Create(&attachment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "file uploaded", "attachment_id": attachment.ID})
}

func DownloadAttachment(c *gin.Context) {
	attachmentID := c.Param("attachment_id")

	var att models.Attachment
	if err := models.DB.First(&att, attachmentID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "attachment not found"})
		return
	}

	c.Data(http.StatusOK, att.ContentType, att.Content)
}
