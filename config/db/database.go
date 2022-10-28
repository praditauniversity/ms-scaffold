package config

import (
	"fmt"
	"ms-scaffold/api/helper"
	"ms-scaffold/api/models/domain"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")
	timeZone := os.Getenv("TZ")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", dbHost, dbUser, dbPass, dbName, dbPort, sslMode, timeZone)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	db.AutoMigrate(&domain.Scaffold{})
	return db
}

func CloseDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	helper.PanicIfError(err)
	dbSQL.Close()
}
