package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"real-time-log-analyze/pkg/models"
)

func InitializeDB() (*gorm.DB, error) {
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	sslmode := os.Getenv("SSL_MODE")
	fmt.Println("Postgres Host:", os.Getenv("POSTGRES_HOST"))
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=%s", host, user, dbname, password, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.LogEntry{})
	return db, nil
}
