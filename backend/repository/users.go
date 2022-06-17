package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

//get all users
func (u *UserRepository) GetAllUsers() ([]User, error) {
	var users []User

	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Fullname, &user.Email, &user.Username, &user.Password, &user.Role, &user.LoggedIn)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, nil
}

//get user's role
func (u *UserRepository) GetUserRole(username string) (*string, error) {
	var role string

	rows, err := u.db.Query("SELECT role FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&role)
		if err != nil {
			return nil, err
		}
	}

	return &role, nil
}

//login
func (u *UserRepository) Login(username string, password string) (*string, error) {
	users, err := u.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			//change loggedin to true
			_, err := u.db.Exec("UPDATE users SET logged_in = ? WHERE username = ?", true, username)
			if err != nil {
				return nil, err
			}
			return &user.Username, nil
		}
	}

	return nil, fmt.Errorf("Login Failed")
}

//insert new user
func (u *UserRepository) Signup(fullname string, email string, username string, password string) error {
	_, err := u.db.Exec("INSERT INTO users (fullname, email, username, password, role, logged_in) VALUES (?, ?, ?, ?, 'user', 'false')", fullname, email, username, password)
	if err != nil {
		return err
	}

	return nil
}

//logout all
func (u *UserRepository) LogoutAll() error {
	_, err := u.db.Exec("UPDATE users SET logged_in = ?", false)
	return err
}
