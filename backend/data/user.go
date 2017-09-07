package data

import (
	"github.com/satori/go.uuid"
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
	Uuid      string
	UserId    int
	LoginTime time.Time
}

// Check if session is valid.
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow("select id, user_id from user_session where uuid = $1", session.Uuid).Scan(&session.Id, &session.UserId)
	if err != nil {
		valid = false
	}
	if session.Id != 0 {
		valid = true
	}
	return
}

//Get user from valid session.
func (session *Session) User() (user User, err error) {
	user = User{}
	err = Db.QueryRow("select id, email, disabled from user_account where id = $1", session.UserId).Scan(&user.Id, &user.Email, &user.Disabled)
	return
}

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

// Create a new session for an existing user.
func (user *User) CreateSession() (session Session, err error) {
	statement := "insert into user_session (uuid, user_id, login_time) values($1, $2, $3) returning id, uuid"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(uuid.NewV4().String(), user.Id, time.Now()).Scan(&session.Id, &session.Uuid)

	return
}

// Retrieve a sesion.
func (user *User) Session() (session Session, err error) {
	session = Session{}

	err = Db.QueryRow("select id, user_id, login_time from user_session where user_id = $1", user.Id).Scan(&session.Id, &session.UserId, &session.LoginTime)

	return
}

// Create a new user
// and save it to DB.
func (user *User) Create() (err error) {
	statement := "insert into user_account(email, password, disabled, created_at) values($1, $2, $3, $4) returning id"
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

// Get user by ID
func UserById(id int) (user User, err error) {
	user = User{}
	statement := "select id, email, disabled from user_account where id = $1"
	err = Db.QueryRow(statement, id).Scan(&user.Id, &user.Email, &user.Disabled)

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

//Checks if user is admin
//func (user *User) IsAdmin() (isAdmin bool, err error) {
//}
