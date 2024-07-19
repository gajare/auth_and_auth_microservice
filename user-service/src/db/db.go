package db

import (
	"fmt"
	"log"
	"user-service/src/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=myapp password=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	DB.AutoMigrate(&model.User{})
	fmt.Println("Database connection successful")
}

func GetDB() *gorm.DB {
	if DB == nil {
		log.Println("Database connection is nil")
	}
	return DB
}
