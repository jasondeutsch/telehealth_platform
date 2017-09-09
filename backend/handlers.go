package main

import (
	"encoding/json"
	"fmt"
	"github.com/jasondeutsch/previ/backend/data"
	"net/http"
)

/**

Patients Resource

**/

func indexPatient(w http.ResponseWriter, r *http.Request) {

	// TODO Check if user is admin.
	// If user is admin then show all patients.
	// TODO Handle errors

	patients, err := data.Patients()

	m := map[string]interface{}{"success": err == nil, "message": "", "data": patients}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func createPatient(w http.ResponseWriter, r *http.Request) {

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

func showPatient(w http.ResponseWriter, r *http.Request) {

	sess, err := session(w, r)
	user, _ := sess.User()

	fmt.Println(user)

	var id string
	err = json.NewDecoder(r.Body).Decode(&id)

	patient, err := data.PatientById(id)

	m := map[string]interface{}{"success": err == nil, "message": "", "data": patient}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)

}

/**

Provider Resource

**/

func createProvider(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)

	var provider *data.Provider
	err := json.NewDecoder(r.Body).Decode(&provider)
	provider.Id = sess.UserId
	fmt.Println(provider)
	err = provider.Create()

	m := map[string]interface{}{"error": err != nil, "message": "", "data": nil}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

func indexProvider(w http.ResponseWriter, r *http.Request) {
	sess, _ := session(w, r)

	fmt.Println(sess)

	providers, err := data.Providers()
	fmt.Println(err)
	fmt.Println(providers)

	m := map[string]interface{}{"error": err != nil, "message": "", "data": providers}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)

}
