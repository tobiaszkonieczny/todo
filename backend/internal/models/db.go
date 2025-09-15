package models

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" port=" + os.Getenv("DB_PORT") +
		" sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	// MIGRATIONS
	// AutoMigrate will create tables, missing foreign keys, constraints, columns and indexes.
	err = DB.AutoMigrate(&User{}, &Category{}, &Task{}, &Attachment{}, Log{})
	if err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}
}
