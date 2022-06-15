package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/rightway.db")
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	sqlStmt :=
		`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fullname VARCHAR(50) NOT NULL,
		email VARCHAR(50) NOT NULL,
		username VARCHAR(64) NOT NULL,
		password VARCHAR(64) NOT NULL,
		role VARCHAR(64) NOT NULL,
		logged_in BOOLEAN NOT NULL
	);
		
	INSERT INTO users (id, fullname, email, username, password, role, logged_in) 
	VALUES (0001, 'admin rightway','admin@admin.com', 'admin', 'admin', 'admin', false);

	CREATE TABLE IF NOT EXISTS fakultas (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		fakultas_name VARCHAR(64) NOT NULL,
		created_at DATETIME NOT NULL
	);

	CREATE TABLE IF NOT EXISTS program_studi (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		prodi_name VARCHAR(64) NOT NULL,
		fakultas_id INTEGER NOT NULL,
		created_at DATETIME NOT NULL
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
