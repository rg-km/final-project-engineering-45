package repository

import "time"

type User struct {
	ID       int64  `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	LoggedIn bool   `db:"logged_in"`
	Token    string `db:"token"`
}

type Fakultas struct {
	ID           int64     `db:"id"`
	FakultasName string    `db:"fakultas_name"`
	CreatedAt    time.Time `db:"created_at"`
}

type ProgramStudi struct {
	ID         int64     `db:"id"`
	ProdiName  string    `db:"prodi_name"`
	FakultasID int64     `db:"fakultas_id"`
	CreatedAt  time.Time `db:"created_at"`
}
