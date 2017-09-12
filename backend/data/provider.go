package data

import (
	"fmt"
	"github.com/lib/pq"
)

type Provider struct {
	Id          int
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	PhoneNumber string   `json:"phone_number"`
	VidyoRoom   string   `json:"vidyo_room"`
	Credential  []string `json:"credential"`
}

// Get Index of Providers
func Providers() (providers []Provider, err error) {
	statement := "select id, first_name, last_name, phone_number, vidyo_room, credential from provider"

	rows, err := Db.Query(statement)
	fmt.Println(rows)

	var p Provider

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.PhoneNumber, &p.VidyoRoom, pq.Array(&p.Credential))

		if err != nil {
			return
		}
		providers = append(providers, p)
	}
	return
}

// Create provider
func (p *Provider) Create(user User) (err error) {
	statement := "insert into provider(id, first_name, last_name, phone_number, credential) values($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Id, p.FirstName, p.LastName, p.PhoneNumber, pq.Array(p.Credential))

	if err == nil {
		err = user.SetRole("provider")
	}

	return

}

// Get Provider By Id
func ProviderById(id int) (p Provider, err error) {
	statement := "select id, first_name, last_name, phone_number, vidyo_room, credential from provider where id=$1"
	stmt, err := Db.Prepare(statement)

	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&p.Id, &p.FirstName, &p.LastName, &p.PhoneNumber, &p.VidyoRoom, pq.Array(&p.Credential))

	return
}

// Get list of patients to which provider is paired.
func (p *Provider) Patients() (patients []Patient, err error) {
	statement := "select id, first_name, last_name from patient where id in (select patient from pairing where provider = $1)"
	stmt, err := Db.Prepare(statement)

	defer stmt.Close()

	rows, err := stmt.Query(p.Id)

	var patient Patient

	for rows.Next() {
		err = rows.Scan(&patient.Id, &patient.FirstName, &patient.LastName)
		if err != nil {
			return
		}
		patients = append(patients, patient)
	}
	return
}

// Check if provider is paired to patient.
// This is used prior to geting patient info.
func (p *Provider) HasPatient(patientId int) (err error) {
	statement := "select count(*) from pairing where provider = $1 and patient = $2"

	_, err = Db.Exec(statement, p.Id, patientId)

	return

}
