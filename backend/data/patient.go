package data

import (
	"time"
)

type Patient struct {
	id         int
	first_name string
	last_name  string
	state      string // timezone may be more to the point?
	country    string
	created_at string
}

// Create a new patient
// and save it to the DB.
func (patient *Patient) Create(user User) (err error) {
	statement := "insert into patient valuse($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	err = stmt.QueryRow(&patient.id, &patient.first_name, &patient.last_name, &patient.state, &patient.country, time.Now()).Scan(&patient.id)

	return
}

// Get next appointment time

// Get their provider/s info

// Get assigned rousources

// View library

// View summary of suggestions

// Messages, send receive to and from paired providers

// Control email notifications

// create/edit health questionnaire

// create/edit food logs

// create/edit MSQ

// upload/share docs
