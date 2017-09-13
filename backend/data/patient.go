package data

import (
	"github.com/lib/pq"
	"time"
)

type Patient struct {
	Id        int
	FirstName string
	LastName  string
	State     string
	Country   string
	CreatedAt string
}

// View all patients
func Patients() (patients []Patient, err error) {

	rows, err := Db.Query("select id, first_name, last_name, state, country from patient")

	var p Patient
	patients = []Patient{}
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.State, &p.Country)
		if err != nil {
			return
		}
		patients = append(patients, p)
	}

	return

}

// Get patient by id
func PatientById(id int) (p Patient, err error) {
	// TODO account for invalid lookup
	// TODO account for authorization
	statement := "select first_name, last_name, state, country from patient where id = $1"
	stmt, err := Db.Prepare(statement)
	err = stmt.QueryRow(id).Scan(&p.FirstName, &p.LastName, &p.State, &p.Country)

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

	_, err = stmt.Exec(&patient.Id, &patient.FirstName, &patient.LastName, &patient.State, &patient.Country, time.Now())

	if err == nil {
		err = user.SetRole("patient")
	}

	return
}

// Get list of providers paired with patient.
func (patient *Patient) Providers() (providers []Provider, err error) {
	statement := "select id, first_name, last_name, credential from provider where id in (select provider from pairing where patient = $1)"
	stmt, err := Db.Prepare(statement)

	defer stmt.Close()

	rows, err := stmt.Query(patient.Id)

	var provider Provider

	for rows.Next() {
		err = rows.Scan(&provider.Id, &provider.FirstName, &provider.LastName, pq.Array(&provider.Credential))
		if err != nil {
			return
		}
		providers = append(providers, provider)
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
