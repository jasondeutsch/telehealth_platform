package data

import "fmt"

//create table patient(
//  id          int primary key references user_account(id),
//  first_name  text not null,
//  last_name   text not null,
//  state       text not null,
//  country     text not null,
//  created_at  timestamp default current_timestamp
//  );

// Pairing, unpair, cancel account.

// View all patients

type Admin struct{}

func (admin *Admin) GetAllPatients() (patients []Patient, err error) {

	rows, err := Db.Query("select id, first_name, last_name, state, country from patient")

	var p Patient
	patients = []Patient{}
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.First_Name, &p.Last_Name, &p.State, &p.Country)
		if err != nil {
			return
		}
		fmt.Println(p)
		patients = append(patients, p)
		fmt.Println(patients)
	}

	return

}

// view patient profile patients surveys, documentation, appointment history

// create provider

// super admin to create admins
