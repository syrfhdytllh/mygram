package database

import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "123456"
	dbPort   = "5433"
	dbname   = "universal"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	db.AutoMigrate(&models.User{}, &models.Photo{}, &models.SocialMedia{}, &models.Comment{})

}

func GetDB() *gorm.DB {

	return db
}
