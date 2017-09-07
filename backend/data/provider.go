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
	// TODO include credentail array in sql query
	statement := "select id, first_name, last_name, phone_number, vidyo_room from provider"

	rows, err := Db.Query(statement)
	fmt.Println(rows)

	var p Provider

	for rows.Next() {
		err = rows.Scan(&p.Id, &p.FirstName, &p.LastName, &p.PhoneNumber, &p.VidyoRoom)
		if err != nil {
			return
		}
		providers = append(providers, p)
	}
	return
}

// Create provider
func (p *Provider) Create() (err error) {
	fmt.Println("Create()")
	statement := "insert into provider(id, first_name, last_name, phone_number, credential) values($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Id, p.FirstName, p.LastName, p.PhoneNumber, pq.Array(p.Credential))

	return

}

func (p *Provider) HasPatient(patientId int) (err error) {
	statement := "select count(*) from pairing where provider = $1 and patient = $2"

	_, err = Db.Exec(statement, p.Id, patientId)

	return

}
