package main

import (
	"github.com/nhatlang19/go-monorepo/services/user/model"
	"github.com/nhatlang19/go-monorepo/services/user/route"
)

func main() {
	db, _ := model.MySqlConnection()
	route.SetupRoutes(db)
}
