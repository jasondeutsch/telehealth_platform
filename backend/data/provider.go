package data

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
	statement := "insert into provider(id, first_name, last_name, phone_number) values($1, $2, $3, $4)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Id, p.FirstName, p.LastName, p.PhoneNumber)

	return

}
