package data

import (
	"time"
)

type Pairing struct {
	PatientId  int
	ProviderId int
	Active     bool
	CreatedAt  time.Time
}

func (p *Pairing) Create() (err error) {
	// Insert IFF pairing does not exist
	statement := "insert into pairing(patient, provider) select $1, $2 from pairing where not exists (select patient, provider from pairing where patient = $1 and provider = $2"
	stmt, err := Db.Prepare(statement)

	defer stmt.Close()

	stmt.Exec(p.PatientId, p.ProviderId)

	return
}

// set active to false
func (p *Pairing) Deactivate() (err error) {
	_, err = Db.Exec("update pairing set active = false where patient=$1 and provider=$2", p.PatientId, p.ProviderId)
	return
}

// set active to true
func (p *Pairing) Activate() (err error) {
	_, err = Db.Exec("update pairing set active = true where patient=$1 and provider=$2", p.PatientId, p.ProviderId)
	return
}
