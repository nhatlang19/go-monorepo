package main

import (
	"github.com/nhatlang19/go-monorepo/services/user/route"
	model "github.com/nhatlang19/go-monorepo/pkg/db"
)

func main() {
	db, _ := model.MySqlConnection()
	route.SetupRoutes(db)
}
