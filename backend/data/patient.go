package data

import (
	"time"
)

type Patient struct {
	Id         int
	First_Name string
	Last_Name  string
	State      string // timezone may be more to the point?
	Country    string
	Created_At string
}

// View all patients

func GetAllPatients() (patients []Patient, err error) {

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
func GetPatientById() {
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

	err = stmt.QueryRow(&patient.Id, &patient.First_Name, &patient.Last_Name, &patient.State, &patient.Country, time.Now()).Scan(&patient.Id)

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
