package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/backend/data"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

/**

Patients Resource

**/

func indexPatient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO Check if user is admin.
	// If user is admin then show all patients.
	// TODO Handle errors

	patients, err := data.Patients()

	m := map[string]interface{}{"success": err == nil, "message": "", "data": patients}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func createPatient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	sess, err := session(w, r)

	user, _ := data.UserById(sess.UserId)

	var patient *data.Patient
	err = json.NewDecoder(r.Body).Decode(&patient)
	fmt.Println(patient)
	err = patient.Create(user)

	m := map[string]interface{}{"success": err == nil, "message": "", "data": user}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)

}

func showPatient(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// TODO Check if user is authorized to view patient.
	// TODO Handle err case of GetPatientById

	id := p.ByName("id")
	patient, err := data.PatientById(id)

	m := map[string]interface{}{"success": err == nil, "message": "", "data": patient}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)

}

/**

Provider Resource

**/

func createProvider(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	_, _ = session(w, r)
}
