package models

import "time"

type Attachment struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TaskID      uint      `json:"task_id"` // foreign key
	FileName    string    `json:"file_name"`
	Content     []byte    `gorm:"type:bytea" json:"-"` // do not expose raw content in JSON
	ContentType string    `json:"content_type"`
	Size        int64     `json:"size"`
	CreatedAt   time.Time `json:"created_at"`
}
