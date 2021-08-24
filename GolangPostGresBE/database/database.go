package database

import (
	"example.com/m/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	connection, err := gorm.Open(postgres.Open("host=localhost user=postgres password=root dbname=go_db port=5432"), &gorm.Config{})
	if err != nil {
		panic("could not connect to DB!")
	}

	DB = connection
	connection.AutoMigrate(&models.User{})

}
