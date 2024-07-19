package main

import (
	"log"
	"net/http"
	"user-service/src/db"
	"user-service/src/router"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db.Init()
	r := router.SetupRouter()

	log.Println("Starting server on :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
