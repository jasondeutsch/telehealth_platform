package data

// view everything about the patient
// create appointments
// summary of suggestions
// create documentation
// send messages
// upload things?

type Provider struct {
	Id          int
	FirstName   string
	LastName    string
	PhoneNumber string
	VidyoRoom   string
	Credential  []string
}

// create provider
func (p *Provider) CreateProvider() (err error) {
	statement := "insert into provider(id, first_name, last_name, phone_number, credential) values($1, $2, $3, $4, $5)"
	stmt, err := Db.Prepare(statement)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = Db.Exec(string(p.Id), p.FirstName, p.LastName, p.PhoneNumber, p.Credential)

	return

}
