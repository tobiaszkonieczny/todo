package models

type User struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"-"`
	Tasks    []Task `gorm:"foreignKey:UserID" json:"tasks"`
}
