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
	Id        int
	UserId    int
	LoginTime time.Time
}

// Create a new session for an existing user.
func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into user_session (user_id, login_time) values($1, $2)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(&session.UserId, time.Now()).Scan(&session.Id)

	return
}

// Retrieve a sesion.
func (user *User) Session() (session Session, err error) {
	session = Session{}

	err = Db.QueryRow("select id, user_id, login_time from user_session where user_id = $1", user.Id).Scan(&session.Id, &session.UserId, &session.LoginTime)

	return
}

// Check if session is valid.

// Delete session
func (session *Session) Delete() (err error) {
	statement := "delete from user_session where id = $1"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(session.Id)

	return
}

// Create a new user
// and save it to DB.
func (user *User) Create() (err error) {
	statement := "insert into user_account (email, password, disabled, created_at) values($1, $2, $3, $4)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Email, Encrypt(user.Password), false, time.Now()).Scan(&user.Id)

	return
}

// Get user by email
func UserByEmail(email string) (user User, err error) {
	user = User{}
	statement := "select id, email, password, disabled, created_at from user_account where email = $1"
	err = Db.QueryRow(statement, email).Scan(&user.Id, &user.Email, &user.Password, &user.Disabled, &user.CreatedAt)

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
