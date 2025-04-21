package config

import (
	"fmt"
	"log"
	"os"

	"r1estate-service/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


user := os.Getenv("DB_USER")
password := os.Getenv("DB_PASSWORD")
host := os.Getenv("DB_HOST")
port := os.Getenv("DB_PORT")
dbname := os.Getenv("DB_NAME")

dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbname)

database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
if err != nil {
	log.Fatal("Failed to connect to database", err)
}

database.AutoMigrate(
    &models.Role{},
    &models.User{},
)


DB = database
fmt.Println("Connected to database")
}