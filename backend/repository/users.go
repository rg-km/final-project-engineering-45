package repository

import (
	"database/sql"
	"fmt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) GetUserByUsername(username string) (*User, error) {
	user := &User{}
	rows, err := u.db.Query("SELECT id, fullname, username, password, role, logged_in FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Fullname, &user.Username, &user.Password, &user.Role, &user.LoggedIn, &user.Token)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

//get all users
func (u *UserRepository) GetAllUsers() ([]User, error) {
	var users []User
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Fullname, &user.Username, &user.Password, &user.Role, &user.LoggedIn, &user.Token)
		if err != nil {
			return nil, err
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
	// err := u.LogoutAll()
	// if err != nil {
	// 	return nil, err
	// }

	users, err := u.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		if user.Username == username && user.Password == password {
			//change logged_in to true
			_, err := u.db.Exec("UPDATE users SET logged_in = ? WHERE username = ?", "true", username)
			if err != nil {
				return nil, err
			}

			return &user.Username, nil
		}
	}

	return nil, fmt.Errorf("username or password is incorrect")
}

//logout
func (u *UserRepository) Logout(username string) error {
	_, err := u.db.Exec("UPDATE users SET logged_in = ? WHERE username = ?", "false", username)
	if err != nil {
		return err
	}

	return nil
}

//insert new user
func (u *UserRepository) SignUp(fullname string, username string, password string, role string, loggedin string) error {
	_, err := u.db.Exec("INSERT INTO users (fullname, username, password, role, logged_in) VALUES (?, ?, ?, ?, ?)", fullname, username, password, role, loggedin)
	if err != nil {
		return err
	}

	return nil
}

//logout all
func (u *UserRepository) LogoutAll() error {
	_, err := u.db.Exec("UPDATE users SET logged_in = ? WHERE logged_in = ?", "false", "true")
	if err != nil {
		return err
	}

	return nil
}
