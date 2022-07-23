package database

import (
	"log"

	"github.com/k-yoshigai/learning-jwt-go/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbErr error

func Connect(connectionString string) {
	Instance, dbErr = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Cannot connect to DB", dbErr)
	}
	log.Println("Connected to DB!")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("DB Migration completed!")
}
