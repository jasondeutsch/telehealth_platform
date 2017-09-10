package data

import (
	"time"
)

type Patient struct {
	Id         int
	First_Name string
	Last_Name  string
	State      string
	Country    string
	Created_At string
}

// View all patients
func Patients() (patients []Patient, err error) {

	rows, err := Db.Query("select id, first_name, last_name, state, country from patient")

	var p Patient
	patients = []Patient{}
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.First_Name, &p.Last_Name, &p.State, &p.Country)
		if err != nil {
			return
		}
		patients = append(patients, p)
	}

	return

}

// Get patient by id
func PatientById(id string) (p Patient, err error) {
	// TODO account for invalid lookup
	// TODO account for authorization
	statement := "select first_name, last_name, state, country from patient where id = $1"
	stmt, err := Db.Prepare(statement)
	err = stmt.QueryRow(id).Scan(&p.First_Name, &p.Last_Name, &p.State, &p.Country)

	return

}

// Create a new patient
// and save it to the DB.
func (patient *Patient) Create(user User) (err error) {
	statement := "insert into patient values($1, $2, $3, $4, $5, $6)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(&patient.Id, &patient.First_Name, &patient.Last_Name, &patient.State, &patient.Country, time.Now())

	if err == nil {
		err = user.SetRole("patient")
	}

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
