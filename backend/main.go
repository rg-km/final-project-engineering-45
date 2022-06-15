package main

import (
	"database/sql"

	"github.com/rg-km/final-project-engineering-45/backend/api"
	"github.com/rg-km/final-project-engineering-45/backend/repository"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "backend/db/rightaway.db")
	if err != nil {
		panic(err)
	}

	usersRepo := repository.NewUserRepository(db)

	mainAPI := api.NewAPI(*usersRepo)
	mainAPI.Start()
}
