package data

// Pairing, unpair, cancel account.

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

// view patient profile patients surveys, documentation, appointment history

//create table provider(
//  id          int primary key references user_account(id),
//  is_admin    bool not null,
//  vidyo_room  text,
//  credential  text
//);

// create provider
//func (admin *Admin) CreateProvider(u User, credential string) (err error) {
//	statement := "insert into provider(id, credential) values($1, $2)"
//	stmt, err := Db.Prepare(statement)
//
//	if err != nil {
//		return
//	}
//
//	defer stmt.Close()
//
//	Db.QueryRow(u.Id)
//
//}

// get provider by id

// super admin to create admins
