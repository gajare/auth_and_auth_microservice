package repository

import (
	"log"
	"user-service/src/db"
	"user-service/src/model"

	"github.com/jinzhu/gorm"
)

func CreateUser(user *model.User) *gorm.DB {
	dbInstance := db.GetDB()
	if dbInstance == nil {
		log.Println("CreateUser: DB instance is nil")
	}
	return dbInstance.Create(user)
}

func GetUserByUsername(username string, user *model.User) *gorm.DB {
	dbInstance := db.GetDB()
	if dbInstance == nil {
		log.Println("GetUserByUsername: DB instance is nil")
	}
	return dbInstance.Where("username = ?", username).First(user)
}

func GetUserByID(id string, user *model.User) *gorm.DB {
	dbInstance := db.GetDB()
	if dbInstance == nil {
		log.Println("GetUserByID: DB instance is nil")
	}
	return dbInstance.Where("id = ?", id).First(user)
}

func UpdateUser(user *model.User) *gorm.DB {
	dbInstance := db.GetDB()
	if dbInstance == nil {
		log.Println("UpdateUser: DB instance is nil")
	}
	return dbInstance.Save(user)
}

func DeleteUser(user *model.User) *gorm.DB {
	dbInstance := db.GetDB()
	if dbInstance == nil {
		log.Println("DeleteUser: DB instance is nil")
	}
	return dbInstance.Delete(user)
}