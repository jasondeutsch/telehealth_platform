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
	statement := "insert into user_account (email, password, disabled, created_at) values($1, $2, $3, $4)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}
	defer stmt.Close()

	// return a row and scan it into User struct
	err = stmt.QueryRow(user.Email, Encrypt(user.Password), false, time.Now()).Scan(&user.Id)

	return
}

// Get user by email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	statement := "select id, email, disabled, created_at from user_account where id=$1"
	err = Db.QueryRow(statement, email).Scan(&user.Id, &user.Email, &user.Disabled, &user.CreatedAt)

	return

}

// Disable user account
func (user *User) Disable() (err error) {
	statement := "update user_account set disabled=true where id=$1"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)

	return
}
