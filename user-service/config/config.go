package config

import (
	"log"
	"user-service/src/db"
)

func Init() {
	db.Init()
	log.Println("init db success...")
}
