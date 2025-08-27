package models

import "time"

type Task struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Title      string    `json:"title"`
	Completed  bool      `json:"completed"`
	UserID     uint      `json:"user_id"` // <-- foreign key
	CategoryID *uint     `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
