package main

import (
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// "github.com/jasondeutsch/previ/backend/data"
/**

Admin

**/

/**

Provider

**/

/**

Patient

**/

// POST
// /logout
func createPatient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, err := session(w, r)

	if err != nil {
		return
	}

	var patient *Patient

	return
}

//create table patient(
//  id          int primary key references user_account(id),
//  first_name  text not null,
//  last_name   text not null,
//  state       text not null,
//  country     text not null,
//  created_at  timestamp default current_timestamp
//  );
//
