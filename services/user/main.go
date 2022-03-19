package main

import (
	"log"

	"github.com/joho/godotenv"
	model "github.com/nhatlang19/go-monorepo/pkg/db"
	"github.com/nhatlang19/go-monorepo/services/user/route"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, _ := model.MySqlConnection()
	route.SetupRoutes(db)
}
