package models

import "time"

type Log struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Method    string    `json:"method"`
	Path      string    `json:"path"`
	IP        string    `json:"ip"`
	UserAgent string    `json:"user_agent"`
	CreatedAt time.Time `json:"created_at"`
}
