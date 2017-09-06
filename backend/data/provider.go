package data

import (
	"fmt"
	"github.com/lib/pq"
)

// view everything about the patient
// create appointments
// summary of suggestions
// create documentation
// send messages
// upload things?

type Provider struct {
	Id          int
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	PhoneNumber string   `json:"phone_number"`
	VidyoRoom   string   `json:"vidyo_room"`
	Credential  []string `json:"credential"`
}

// create provider
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
