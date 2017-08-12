package data

import (
	"time"
)

type User struct {
	Id        int
	Email     string
	Password  string
	Disabled  bool
	CreatedAt time.Time
}

type Session struct {
	Id           int
	UserId       int
	LoginTime    time.Time
	LastSeenTime time.Time
}

// Create a new session for an existing user.

// Retrieve a sesion.

// Check if session is valid.

// Delete session

// Create a new user
// and save it to DB.
func (user *User) Create() (err error) {
	statement := "insert into user_account (email, password) values($1, $2)"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	// return a row and scan it into User struct
	err = stmt.QueryRow(user.Email, user.Disabled, Encrypt(user.Password), user.CreatedAt).Scan(&user.Id)
	return
}

// Get user by email

// Disable user account
